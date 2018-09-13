package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port          string
	PublicKeyPath string
	DBConfig            DBConfig `toml:"database"`
}

var configure Config

func LoadConfig(filePath string) error {
	if _, err := toml.DecodeFile(filePath, &configure); err != nil {
		return errors.New("Failed to load config : " + err.Error())
	}
	return nil
}

func GetConfig() Config {
	return configure
}
