package errs

import "net/http"

// APIErrors struct
type APIErrors struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// AsMessage funv
func (e APIErrors) AsMessage() *APIErrors {
	return &APIErrors{
		Message: e.Message,
	}
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
