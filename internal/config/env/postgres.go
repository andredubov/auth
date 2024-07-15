package env

import (
	"errors"
	"os"

	"github.com/andredubov/auth/internal/config"
)

const (
	dsnEnvName = "PG_DSN"
)

type postgresConfig struct {
	dsn string
}

// NewPostgresConfig returns an intance of postgresConfig struct
func NewPostgresConfig() (config.PostgresConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("pg dsn not found")
	}

	return &postgresConfig{
		dsn: dsn,
	}, nil
}

// DSN returns postgres database connecton string
func (cfg *postgresConfig) DSN() string {
	return cfg.dsn
}
