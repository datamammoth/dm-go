package datamammoth

import "fmt"

// Error represents an API error response from the DataMammoth API.
type Error struct {
	// StatusCode is the HTTP status code.
	StatusCode int

	// RequestID is the unique request identifier from the API.
	RequestID string

	// Errors contains the individual error objects.
	Errors []APIError

	// RawBody contains the raw JSON response body (for debugging).
	RawBody string
}

// Error implements the error interface.
func (e *Error) Error() string {
	if len(e.Errors) == 0 {
		return fmt.Sprintf("datamammoth: API error (status %d, request %s)", e.StatusCode, e.RequestID)
	}
	first := e.Errors[0]
	if len(e.Errors) == 1 {
		return fmt.Sprintf("datamammoth: %s — %s (status %d, request %s)",
			first.Code, first.Message, e.StatusCode, e.RequestID)
	}
	return fmt.Sprintf("datamammoth: %s — %s (+%d more errors) (status %d, request %s)",
		first.Code, first.Message, len(e.Errors)-1, e.StatusCode, e.RequestID)
}

// IsNotFound reports whether the error is a 404 Not Found.
func IsNotFound(err error) bool {
	e, ok := err.(*Error)
	return ok && e.StatusCode == 404
}

// IsRateLimited reports whether the error is a 429 Too Many Requests.
func IsRateLimited(err error) bool {
	e, ok := err.(*Error)
	return ok && e.StatusCode == 429
}

// IsValidation reports whether the error is a 400 Validation Failed.
func IsValidation(err error) bool {
	e, ok := err.(*Error)
	return ok && e.StatusCode == 400
}

// IsAuthError reports whether the error is a 401 Authentication error.
func IsAuthError(err error) bool {
	e, ok := err.(*Error)
	return ok && e.StatusCode == 401
}

// IsPermission reports whether the error is a 403 Permission Denied.
func IsPermission(err error) bool {
	e, ok := err.(*Error)
	return ok && e.StatusCode == 403
}

// IsConflict reports whether the error is a 409 Conflict.
func IsConflict(err error) bool {
	e, ok := err.(*Error)
	return ok && e.StatusCode == 409
}

// IsServerError reports whether the error is a 5xx server error.
func IsServerError(err error) bool {
	e, ok := err.(*Error)
	return ok && e.StatusCode >= 500
}
