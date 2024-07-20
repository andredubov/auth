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
	// ErrEmptyConfigFilePath is error config path is empty
	ErrEmptyConfigFilePath = errors.New("config path is empty")
	// ErrEmptyConfigFilePath is error config file does't exist
	ErrConfigFileDoesNotExist = errors.New("config file does't exist")
)

// GRPCConfig interface
type GRPCConfig interface {
	Address() string
}

// PostgresConfig interface
type PostgresConfig interface {
	DSN() string
}

// AuthConfing interface
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
