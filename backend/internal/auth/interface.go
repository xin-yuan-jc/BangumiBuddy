package auth

import (
	"context"

	"github.com/MangataL/BangumiBuddy/pkg/errs"
)

var (
	ErrUsernameOrPasswordError = errs.NewUnauthorized("用户名或密码错误")
)

// Authenticator 鉴权接口
type Authenticator interface {
	// Authorize 用户密码方式进行授权
	Authorize(ctx context.Context, username, password string) (Credentials, error)
	// ChangePassword 改变密码
	ChangePassword(ctx context.Context, token, newPassword string) error
	// CheckAccessToken 检查token是否有效
	CheckAccessToken(ctx context.Context, token string) error
	// RefreshCredentials 刷新凭证
	RefreshCredentials(ctx context.Context, refreshToken string) (Credentials, error)
}
