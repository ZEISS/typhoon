package apis

import (
	"fmt"
	"net/http"
)

// Error represents an error that can be returned by the API.
func (e Error) Error() error {
	return fmt.Errorf("code: %d, message: %s", e.Code, e.Message)
}

// Unimplemented is a struct that implements the StrictServerInterface interface
func NotImplemented(message string, args ...interface{}) Error {
	return Error{
		Code:    http.StatusNotImplemented,
		Message: fmt.Sprintf(message, args...),
	}
}
