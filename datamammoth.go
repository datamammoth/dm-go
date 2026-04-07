// Package datamammoth provides an idiomatic Go client for the DataMammoth API v2.
//
// Usage:
//
//	client := datamammoth.NewClient("dm_your_api_key")
//	servers, _, err := client.Servers.List(ctx, nil)
package datamammoth

import (
	"net/http"
	"time"
)

// Client is the top-level DataMammoth API client. Use NewClient to create one.
type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client

	Servers   *ServersService
	Products  *ProductsService
	Billing   *BillingService
	Support   *SupportService
	Account   *AccountService
	Admin     *AdminService
	Affiliate *AffiliateService
	Webhooks  *WebhooksService
	Tasks     *TasksService
	Zones     *ZonesService
}

// Option configures a Client.
type Option func(*Client)

// NewClient creates a new DataMammoth API client.
//
// The apiKey is required and should be a valid DataMammoth API key
// (e.g. "dm_abc123..."). Use functional options to override defaults:
//
//	client := datamammoth.NewClient("dm_key",
//	    datamammoth.WithBaseURL("https://staging.datamammoth.com/api/v2"),
//	    datamammoth.WithTimeout(60 * time.Second),
//	)
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:  apiKey,
		baseURL: "https://app.datamammoth.com/api/v2",
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Servers = &ServersService{client: c}
	c.Products = &ProductsService{client: c}
	c.Billing = &BillingService{client: c}
	c.Support = &SupportService{client: c}
	c.Account = &AccountService{client: c}
	c.Admin = &AdminService{client: c}
	c.Affiliate = &AffiliateService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	c.Tasks = &TasksService{client: c}
	c.Zones = &ZonesService{client: c}

	return c
}

// WithBaseURL overrides the default API base URL.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithTimeout overrides the default HTTP client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.http.Timeout = d
	}
}

// WithHTTPClient replaces the default HTTP client entirely.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.http = hc
	}
}
