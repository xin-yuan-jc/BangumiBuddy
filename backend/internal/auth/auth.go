package auth

import (
	"context"
	"time"

	"github.com/MangataL/BangumiBuddy/internal/config"
	"github.com/MangataL/BangumiBuddy/pkg/errs"
	"github.com/MangataL/BangumiBuddy/pkg/log"
)

//go:generate mockgen -destination auth_mock.go -source $GOFILE -package $GOPACKAGE

// New 生成鉴权器
func New(dep Dependencies) Authenticator {
	return &authenticator{
		config:        dep.Config,
		cipher:        dep.Cipher,
		tokenOperator: dep.TokenOperator,
	}
}

// Dependencies 依赖
type Dependencies struct {
	config.Config
	Cipher
	TokenOperator
}

// Cipher is the interface for token operation
type Cipher interface {
	GenerateKey(ctx context.Context) (string, error)
	Encrypt(ctx context.Context, token, text string) (string, error)
	Check(ctx context.Context, token, text, cipher string) error
}

// TokenOperator token操作器
type TokenOperator interface {
	Generate(ctx context.Context, username, key string, expireAt time.Time) (string, error)
	Check(ctx context.Context, key, token string) error
}

// authenticator
type authenticator struct {
	config        config.Config
	cipher        Cipher
	tokenOperator TokenOperator
}

const (
	defaultAccessTokenExpire  = time.Hour * 24
	defaultRefreshTokenExpire = time.Hour * 24 * 7
)

func (a *authenticator) Authorize(ctx context.Context, username, password string) (Credentials, error) {
	token := a.getToken()
	if err := a.validateLogin(ctx, username, password, token); err != nil {
		return Credentials{}, err
	}
	credentials, err := a.generateCredentials(ctx, username, token)
	if err != nil {
		return Credentials{}, err
	}
	return credentials, nil
}

func (a *authenticator) validateLogin(ctx context.Context, username, password, token string) error {
	if username == "" {
		return errs.NewUnauthorized("用户名不能为空")
	}
	if username != a.getUsername() {
		return ErrUsernameOrPasswordError
	}
	pw := a.getPassword()
	if err := a.cipher.Check(ctx, token, password, pw); err != nil {
		log.Error(ctx, "检查密码错误: ", err.Error())
		return ErrUsernameOrPasswordError
	}
	return nil
}

func (a *authenticator) getUsername() string {
	username, _ := a.config.GetUsername()
	return username
}

func (a *authenticator) getPassword() string {
	password, _ := a.config.GetPassword()
	return password
}

func (a *authenticator) getToken() string {
	token, _ := a.config.GetToken()
	return token
}

func (a *authenticator) generateCredentials(ctx context.Context, username, token string) (Credentials, error) {
	accessToken, err := a.tokenOperator.Generate(ctx, username, token, time.Now().Add(defaultAccessTokenExpire))
	if err != nil {
		return Credentials{}, err
	}
	refreshToken, err := a.tokenOperator.Generate(ctx, username, token, time.Now().Add(defaultRefreshTokenExpire))
	if err != nil {
		return Credentials{}, err
	}
	return Credentials{
		AccessToken:  accessToken,
		TokenType:    "bearer",
		RefreshToken: refreshToken,
	}, nil
}

func (a *authenticator) ChangePassword(ctx context.Context, accessToken, newPassword string) error {
	// TODO implement me
	panic("implement me")
}

func (a *authenticator) CheckAccessToken(ctx context.Context, accessToken string) error {
	// TODO implement me
	panic("implement me")
}

func (a *authenticator) RefreshCredentials(ctx context.Context, refreshToken string) (Credentials, error) {
	// TODO implement me
	panic("implement me")
}
