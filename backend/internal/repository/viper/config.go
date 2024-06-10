package viper

import (
	"github.com/spf13/viper"
)

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
