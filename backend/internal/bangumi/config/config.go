package config

type TMDB interface {
	GetAPIToken() (string, error)
	SetAPIToken(token string) error
}
