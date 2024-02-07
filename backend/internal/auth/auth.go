package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/MangataL/BangumiBuddy/pkg/log"
)

//go:generate mockgen -destination auth_mock.go -source $GOFILE -package $GOPACKAGE

// UserRepository is the interface for user data
type UserRepository interface {
	GetUser(ctx context.Context) (User, error)
	ApplyUser(ctx context.Context, user User) error
	GetUserCookie(ctx context.Context, cookie string) (UserCookie, error)
	SetCookie(ctx context.Context, cookie UserCookie) error
	DeleteCookie(ctx context.Context, username string) error
}

// UserCookie is the struct for user cookie
type UserCookie struct {
	Username string
	Cookie   string
	ExpireAt time.Time
}

// authenticator based on Cookie-Session Authentication
type authenticator struct {
	userRepo UserRepository
}

func (a *authenticator) Login(ctx context.Context, username, password string) (http.Cookie, error) {
	user, err := a.userRepo.GetUser(ctx)
	if err != nil {
		return http.Cookie{}, err
	}
	if user.Username != username || user.Password != password {
		return http.Cookie{}, errors.New("用户名或密码错误")
	}
	cookie := a.newCookie()
	if err := a.userRepo.SetCookie(ctx, UserCookie{
		Username: username,
		Cookie:   cookie.Value,
		ExpireAt: cookie.Expires,
	}); err != nil {
		return http.Cookie{}, err
	}
	return cookie, nil
}

func (a *authenticator) Logout(ctx context.Context, username string) error {
	return a.userRepo.DeleteCookie(ctx, username)
}

func (a *authenticator) UpdateUser(ctx context.Context, username, password string) error {
	return a.userRepo.ApplyUser(ctx, User{
		Username: username,
		Password: password,
	})
}

// User is the struct for user data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	// CookieName is the name of the cookie
	CookieName = "bangumi_session"
)

func (a *authenticator) CheckCookie(r *http.Request) (http.Cookie, error) {
	cookie, err := r.Cookie(CookieName)
	if err != nil {
		return http.Cookie{}, err
	}
	userCookie, err := a.userRepo.GetUserCookie(r.Context(), cookie.Value)
	if err != nil {
		return http.Cookie{}, err
	}
	now := time.Now()
	if userCookie.ExpireAt.Before(now) {
		return http.Cookie{}, errors.New("cookie过期，请重新登录")
	}
	if userCookie.ExpireAt.Sub(now) < 10*time.Minute {
		newCookie := a.newCookie()
		err = a.userRepo.SetCookie(r.Context(), UserCookie{
			Username: userCookie.Username,
			Cookie:   newCookie.Value,
			ExpireAt: newCookie.Expires,
		})
		if err != nil {
			log.Errorf(r.Context(), "set cookie failed %s", err)
		}
		return newCookie, nil
	}
	return http.Cookie{}, nil
}

func (a *authenticator) newCookie() http.Cookie {
	return http.Cookie{
		Name:    CookieName,
		Value:   uuid.NewString(),
		MaxAge:  24 * 60 * 60,
		Expires: time.Now().Add(24 * time.Hour),
	}
}

// NewAuth returns a new authenticator
func NewAuth(userRepo UserRepository) Authenticator {
	return &authenticator{
		userRepo: userRepo,
	}
}
