package auth

import (
	"context"
	"net/http"
)

// Authenticator is the interface for authentication
type Authenticator interface {
	Login(ctx context.Context, username, password string) (http.Cookie, error)
	Logout(ctx context.Context, username string) error
	UpdateUser(ctx context.Context, username, password string) error
	CheckCookie(r *http.Request) (http.Cookie, error)
}
