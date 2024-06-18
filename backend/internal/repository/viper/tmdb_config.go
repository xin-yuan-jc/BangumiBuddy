package viper

import (
	"github.com/MangataL/BangumiBuddy/internal/bangumi/config"
)

var _ config.TMDB = (*Repo)(nil)

func (r *Repo) GetAPIToken() (string, error) {
	return r.file.GetString("tmdb.api_token"), nil
}

func (r *Repo) SetAPIToken(token string) error {
	r.file.Set("tmdb.api_token", token)
	return r.file.WriteConfig()
}
