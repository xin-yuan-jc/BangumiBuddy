package errs

import (
	"net/http"

	"github.com/pkg/errors"
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

func ParseError(err error) (int, string) {
	cause := errors.Cause(err)
	if e := (&Error{}); errors.As(cause, &e) {
		return e.HTTPCode, e.Error()
	}
	return http.StatusInternalServerError, err.Error()
}
