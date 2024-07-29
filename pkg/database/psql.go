package database

import (
	"context"

	"github.com/andredubov/auth/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// NewPostgresConnection returns an instance for connection to the postgres database
func NewPostgresConnection(ctx context.Context, cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(ctx, cfg.DSN())
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
