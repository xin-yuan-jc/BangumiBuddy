package config

//go:generate mockgen -destination config_mock.go -source $GOFILE -package $GOPACKAGE

// Config is the configuration for BangumiBuddy
type Config interface {
	GetUsername() (string, error)
	SetUsername(username string) error
	GetPassword() (string, error)
	SetPassword(password string) error
	GetToken() (string, error)
	SetToken(token string) error
}
