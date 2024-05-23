package errs

import (
	"net/http"
)

// Error 异常
type Error struct {
	HTTPCode int
	Message  string
}

func (e *Error) Error() string {
	return e.Message
}

func NewForbidden(message string) *Error {
	return &Error{
		HTTPCode: http.StatusForbidden,
		Message:  message,
	}
}

func NewUnauthorized(message string) *Error {
	return &Error{
		HTTPCode: http.StatusUnauthorized,
		Message:  message,
	}
}
