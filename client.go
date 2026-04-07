package datamammoth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// apiResponse is the raw envelope returned by the DataMammoth API.
type apiResponse struct {
	Data   json.RawMessage `json:"data"`
	Meta   Meta            `json:"meta"`
	Errors []APIError      `json:"errors"`
	Links  Links           `json:"_links,omitempty"`
}

// do executes an HTTP request against the API and decodes the response into dest.
// If dest is nil, the response body is not decoded (useful for 204 No Content).
func (c *Client) do(ctx context.Context, method, path string, body interface{}, dest interface{}) (*Meta, error) {
	u := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("datamammoth: failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("datamammoth: failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", "datamammoth-go/0.1.0")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("datamammoth: request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("datamammoth: failed to read response body: %w", err)
	}

	// Handle non-JSON responses (e.g., 204 No Content)
	if len(respBody) == 0 {
		if resp.StatusCode >= 400 {
			return nil, &Error{
				StatusCode: resp.StatusCode,
			}
		}
		return &Meta{}, nil
	}

	var envelope apiResponse
	if err := json.Unmarshal(respBody, &envelope); err != nil {
		// If the response is not valid JSON, wrap it in an error
		if resp.StatusCode >= 400 {
			return nil, &Error{
				StatusCode: resp.StatusCode,
				RawBody:    string(respBody),
			}
		}
		return nil, fmt.Errorf("datamammoth: failed to decode response: %w", err)
	}

	// Check for API errors
	if resp.StatusCode >= 400 || len(envelope.Errors) > 0 {
		return &envelope.Meta, &Error{
			StatusCode: resp.StatusCode,
			RequestID:  envelope.Meta.RequestID,
			Errors:     envelope.Errors,
			RawBody:    string(respBody),
		}
	}

	// Decode the data field into dest
	if dest != nil && envelope.Data != nil {
		if err := json.Unmarshal(envelope.Data, dest); err != nil {
			return &envelope.Meta, fmt.Errorf("datamammoth: failed to decode data: %w", err)
		}
	}

	return &envelope.Meta, nil
}

// doList executes a paginated list request and returns both the data and pagination info.
func (c *Client) doList(ctx context.Context, path string, opts *ListOptions, dest interface{}) (*Pagination, error) {
	u := c.buildListURL(path, opts)

	meta, err := c.do(ctx, http.MethodGet, u, nil, dest)
	if err != nil {
		return nil, err
	}

	if meta != nil && meta.Pagination != nil {
		return meta.Pagination, nil
	}

	return nil, nil
}

// buildListURL constructs a URL with pagination, sort, search, and filter query params.
func (c *Client) buildListURL(path string, opts *ListOptions) string {
	if opts == nil {
		return path
	}

	params := url.Values{}
	if opts.Page > 0 {
		params.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.PerPage > 0 {
		params.Set("per_page", strconv.Itoa(opts.PerPage))
	}
	if opts.Sort != "" {
		params.Set("sort", opts.Sort)
	}
	if opts.Search != "" {
		params.Set("search", opts.Search)
	}
	for k, v := range opts.Filter {
		params.Set("filter["+k+"]", v)
	}

	q := params.Encode()
	if q == "" {
		return path
	}

	if strings.Contains(path, "?") {
		return path + "&" + q
	}
	return path + "?" + q
}

// get is a convenience method for GET requests.
func (c *Client) get(ctx context.Context, path string, dest interface{}) (*Meta, error) {
	return c.do(ctx, http.MethodGet, path, nil, dest)
}

// post is a convenience method for POST requests.
func (c *Client) post(ctx context.Context, path string, body interface{}, dest interface{}) (*Meta, error) {
	return c.do(ctx, http.MethodPost, path, body, dest)
}

// patch is a convenience method for PATCH requests.
func (c *Client) patch(ctx context.Context, path string, body interface{}, dest interface{}) (*Meta, error) {
	return c.do(ctx, http.MethodPatch, path, body, dest)
}

// put is a convenience method for PUT requests.
func (c *Client) put(ctx context.Context, path string, body interface{}, dest interface{}) (*Meta, error) {
	return c.do(ctx, http.MethodPut, path, body, dest)
}

// del is a convenience method for DELETE requests.
func (c *Client) del(ctx context.Context, path string, dest interface{}) (*Meta, error) {
	return c.do(ctx, http.MethodDelete, path, nil, dest)
}
