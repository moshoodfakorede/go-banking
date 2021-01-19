package errs

import "net/http"

// APIErrors struct
type APIErrors struct {
	Code    int
	Message string
}

// NewNotFoundError method
func NewNotFoundError(message string) *APIErrors {
	return &APIErrors{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewUnexpectedError method
func NewUnexpectedError(message string) *APIErrors {
	return &APIErrors{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
