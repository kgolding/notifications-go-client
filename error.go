package notify

import (
	"fmt"
)

// APIError is a custom error-type struct adjusted for the API usage.
type APIError struct {
	Message    string
	StatusCode int
	Errors     []Error
}

// Error method is here to return a top level error message as well as define
// our new error type.
func (e *APIError) Error() string {
	m := fmt.Sprintf("%s [Code: %d]", e.Message, e.StatusCode)
	if e != nil && len(e.Errors) > 0 {
		var s string
		for _, e := range e.Errors {
			s += ", " + e.Error + ": " + e.Message
		}
		m += ": " + s[1:]
	}
	return m
}

// Error may be returned by the API.
type Error struct {
	Error   string
	Message string
}
