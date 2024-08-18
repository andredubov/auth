package app

import (
	"context"
	"log"

	server "github.com/andredubov/auth/internal/api/auth/v1"
	"github.com/andredubov/auth/internal/cache"
	"github.com/andredubov/auth/internal/cache/user/redis"
	"github.com/andredubov/auth/internal/repository"
	postgres "github.com/andredubov/auth/internal/repository/user/postgres"
	"github.com/andredubov/auth/internal/service"
	"github.com/andredubov/auth/internal/service/user"
	cacheCl "github.com/andredubov/golibs/pkg/client/cache"
	redisClient "github.com/andredubov/golibs/pkg/client/cache/redis"
	"github.com/andredubov/golibs/pkg/client/database"
	postgresClient "github.com/andredubov/golibs/pkg/client/database/postgres"
	"github.com/andredubov/golibs/pkg/client/database/transaction"
	"github.com/andredubov/golibs/pkg/closer"
	"github.com/andredubov/golibs/pkg/config"
	"github.com/andredubov/golibs/pkg/config/env"
	"github.com/andredubov/golibs/pkg/hasher"
)

type serviceProvider struct {
	postgresConfig       config.PostgresConfig
	authConfig           config.AuthConfing
	grpcConfig           config.GRPCConfig
	redisConfig          config.RedisConfig
	passwordHasher       hasher.PasswordHasher
	databaseClient       database.Client
	databaseTxManager    database.TxManager
	usersRepository      repository.Users
	cacheClient          cacheCl.Client
	usersCache           cache.Users
	usersService         service.Users
	serverImplementation *server.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// AuthConfig loads auth config from appropriate enviroment variables
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

// PostgresConfig loads postges config from appropriate enviroment variables
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

// RedisConfig loads postges config from appropriate enviroment variables
func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %s", err.Error())
		}

		s.redisConfig = cfg
	}
	return s.redisConfig
}

// GRPCConfig loads grpc config from appropriate enviroment variables
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

// PasswordHasher creates an instance of password hasher
func (s *serviceProvider) PasswordHasher() hasher.PasswordHasher {
	if s.passwordHasher == nil {
		salt := s.AuthConfig().PasswordSalt()
		s.passwordHasher = hasher.NewSHA256Hasher(salt)
	}

	return s.passwordHasher
}

// DatabaseClient creates an instance of database client
func (s *serviceProvider) DatabaseClient(ctx context.Context) database.Client {
	if s.databaseClient == nil {
		dbClient, err := postgresClient.New(ctx, s.PostgresConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		if err := dbClient.Database().Ping(ctx); err != nil {
			log.Fatalf("database ping error: %v", err)
		}

		closer.Add(func() error {
			dbClient.Database().Close()
			return nil
		})

		s.databaseClient = dbClient
	}

	return s.databaseClient
}

// CacheClient creates an instance of cache client
func (s *serviceProvider) CacheClient(ctx context.Context) cacheCl.Client {
	if s.cacheClient == nil {
		client, err := redisClient.New(ctx, s.RedisConfig())
		if err != nil {
			log.Fatalf("failed to connect to cache: %v", err)
		}

		if err := client.Cache().Ping(ctx); err != nil {
			log.Fatalf("cache ping error: %v", err)
		}

		closer.Add(func() error {
			return client.Cache().Close()
		})

		s.cacheClient = client
	}

	return s.cacheClient
}

// TxManager creates an instance of transaction managet
func (s *serviceProvider) TxManager(ctx context.Context) database.TxManager {
	if s.databaseTxManager == nil {
		db := s.DatabaseClient(ctx).Database()
		s.databaseTxManager = transaction.NewTransactionManager(db)
	}

	return s.databaseTxManager
}

// UsersRepository creates an instance of users repository
func (s *serviceProvider) UsersRepository(ctx context.Context) repository.Users {
	if s.usersRepository == nil {
		dbClient := s.DatabaseClient(ctx)
		s.usersRepository = postgres.NewUsersRepository(dbClient)
	}

	return s.usersRepository
}

// UsersCache creates an instance of users repository
func (s *serviceProvider) UsersCache(ctx context.Context) cache.Users {
	if s.usersCache == nil {
		cacheClient := s.CacheClient(ctx)
		s.usersCache = redis.NewUsersCache(cacheClient)
	}

	return s.usersCache
}

// UsersService creates an instance of users service
func (s *serviceProvider) UsersService(ctx context.Context) service.Users {
	if s.usersService == nil {
		s.usersService = user.NewService(
			s.UsersRepository(ctx),
			s.PasswordHasher(),
			s.UsersCache(ctx),
			s.TxManager(ctx),
		)
	}

	return s.usersService
}

// ServerImplementation creates an instance of grpc server implementation
func (s *serviceProvider) ServerImplementation(ctx context.Context) *server.Implementation {
	if s.serverImplementation == nil {
		usersService := s.UsersService(ctx)
		s.serverImplementation = server.NewImplementation(usersService)
	}

	return s.serverImplementation
}
