package pbkdf2

import (
	"crypto/rand"
	"encoding/base64"
)

// Operator 使用 pbkdf2 实现的加密解密器
type Operator struct{}

// NewOperator 生成 pbkdf2 实现的加密解密工具
func NewOperator() *Operator {
	return &Operator{}
}

// GenerateToken 生成 token
func (o *Operator) GenerateToken() (string, error) {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func (o *Operator) Encrypt(token string) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (o *Operator) Check(token, cipher string) (bool, error) {
	// TODO implement me
	panic("implement me")
}
