package viper

import (
	"github.com/spf13/viper"

	"github.com/MangataL/BangumiBuddy/internal/config"
)

var _ config.Config = (*Repo)(nil)

// Repo is the repository
type Repo struct {
	file *viper.Viper
}

func NewRepo(path string) (*Repo, error) {
	file := viper.New()
	file.SetConfigFile(path)
	if err := file.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Repo{
		file: file,
	}, nil
}

func (r *Repo) GetUsername() (string, error) {
	return r.file.GetString("user.username"), nil
}

func (r *Repo) SetUsername(username string) error {
	r.file.Set("user.username", username)
	return r.file.WriteConfig()
}

func (r *Repo) GetPassword() (string, error) {
	return r.file.GetString("user.password"), nil
}

func (r *Repo) SetPassword(password string) error {
	r.file.Set("user.password", password)
	return r.file.WriteConfig()
}

func (r *Repo) GetToken() (string, error) {
	return r.file.GetString("user.token"), nil
}

func (r *Repo) SetToken(token string) error {
	r.file.Set("user.token", token)
	return r.file.WriteConfig()
}
