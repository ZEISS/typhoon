package apis

import (
	"fmt"
	"net/http"
)

// Error represents an error that can be returned by the API.
func (e Error) Error() error {
	return fmt.Errorf("code: %d, message: %s", e.Code, e.Message)
}

// NotImplemented is a helper function to create a new Error with a 501 status code.
func NotImplemented(message string, args ...interface{}) Error {
	return Error{
		Code:    http.StatusNotImplemented,
		Message: fmt.Sprintf(message, args...),
	}
}

// ErrorNotFound is a helper function to create a new Error with a 404 status code.
func ErrorNotFound(message string, args ...interface{}) Error {
	return Error{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf(message, args...),
	}
}

// ErrorBadRequest is a helper function to create a new Error with a 400 status code.
func ErrorBadRequest(message string, args ...interface{}) Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: fmt.Sprintf(message, args...),
	}
}

// ErrorUnauthorized is a helper function to create a new Error with a 401 status code.
func ErrorUnauthorized(message string, args ...interface{}) Error {
	return Error{
		Code:    http.StatusUnauthorized,
		Message: fmt.Sprintf(message, args...),
	}
}
