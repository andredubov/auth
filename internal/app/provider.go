package app

import (
	"context"
	"log"

	server "github.com/andredubov/auth/internal/api/auth/v1"
	"github.com/andredubov/auth/internal/closer"
	"github.com/andredubov/auth/internal/config"
	"github.com/andredubov/auth/internal/config/env"
	"github.com/andredubov/auth/internal/repository"
	postgres "github.com/andredubov/auth/internal/repository/postgres/user"
	"github.com/andredubov/auth/internal/service"
	"github.com/andredubov/auth/internal/service/user"
	"github.com/andredubov/auth/pkg/hasher"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	postgresConfig       config.PostgresConfig
	authConfig           config.AuthConfing
	grpcConfig           config.GRPCConfig
	passwordHasher       hasher.PasswordHasher
	postgresPool         *pgxpool.Pool
	usersRepository      repository.Users
	usersService         service.Users
	serverImplementation *server.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) AuthConfig() config.AuthConfing {
	if s.authConfig == nil {
		cfg, err := env.NewAuthConfig()
		if err != nil {
			log.Fatalf("failed to get auth config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) PostgresConfig() config.PostgresConfig {
	if s.postgresConfig == nil {
		cfg, err := env.NewPostgresConfig()
		if err != nil {
			log.Fatalf("failed to get postgres config: %s", err.Error())
		}

		s.postgresConfig = cfg
	}

	return s.postgresConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PasswordHasher() hasher.PasswordHasher {
	if s.passwordHasher == nil {
		salt := s.AuthConfig().PasswordSalt()
		s.passwordHasher = hasher.NewSHA256Hasher(salt)
	}

	return s.passwordHasher
}

func (s *serviceProvider) PostgresPool(ctx context.Context) *pgxpool.Pool {
	if s.postgresPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PostgresConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		if err := pool.Ping(ctx); err != nil {
			log.Fatalf("database ping error: %v", err)
		}

		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.postgresPool = pool
	}

	return s.postgresPool
}

func (s *serviceProvider) UsersRepository(ctx context.Context) repository.Users {
	if s.usersRepository == nil {
		pool := s.PostgresPool(ctx)
		s.usersRepository = postgres.NewUsersRepository(pool)
	}

	return s.usersRepository
}

func (s *serviceProvider) UsersService(ctx context.Context) service.Users {
	if s.usersService == nil {
		usersRepository := s.UsersRepository(ctx)
		passwordHasher := s.PasswordHasher()
		s.usersService = user.NewService(usersRepository, passwordHasher)
	}

	return s.usersService
}

func (s *serviceProvider) ServerImplementation(ctx context.Context) *server.Implementation {
	if s.serverImplementation == nil {
		usersService := s.UsersService(ctx)
		s.serverImplementation = server.NewImplementation(usersService)
	}

	return s.serverImplementation
}
