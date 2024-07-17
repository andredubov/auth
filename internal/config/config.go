package config

import (
	"errors"
	"flag"
	"os"

	"github.com/joho/godotenv"
)

const (
	ConfigPathFlagName   = "config-path"
	ConfigPathEnvVarName = "CONFIG_PATH"
	EmptyString          = ""
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

// Load loads env variables from env-file
func Load() error {
	configPath := fetchConfigPath()
	if configPath == EmptyString {
		return ErrEmptyConfigFilePath
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return ErrConfigFileDoesNotExist
	}

	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

// fetchConfigPath fetches config path from command line flag or enviroment variable
// priority: flag > env > default
// default value is empty string
func fetchConfigPath() string {
	var result string

	if flag.Lookup(ConfigPathFlagName) == nil {
		flag.StringVar(&result, ConfigPathFlagName, EmptyString, "path to env file")
	}
	flag.Parse()

	if result == EmptyString {
		result = os.Getenv(ConfigPathEnvVarName)
	}

	return result
}
