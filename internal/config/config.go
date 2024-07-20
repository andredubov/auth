package config

import (
	"errors"
	"flag"

	"github.com/joho/godotenv"
)

const (
	ConfigPathFlagName = "config-path"
	EmptyString        = ""
)

var (
	ErrEmptyConfigFilePath    = errors.New("config path is empty")
	ErrConfigFileDoesNotExist = errors.New("config file does't exist")
)

type GRPCConfig interface {
	Address() string
}

type PostgresConfig interface {
	DSN() string
}

type AuthConfing interface {
	PasswordSalt() string
}

// Load enviriment variables from *.env file
func Load() error {
	var configPath string

	if flag.Lookup(ConfigPathFlagName) == nil {
		flag.StringVar(&configPath, ConfigPathFlagName, EmptyString, "path to env file")
	}
	flag.Parse()

	if configPath == EmptyString {
		return nil
	}

	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}
