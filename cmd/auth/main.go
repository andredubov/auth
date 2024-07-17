package main

import (
	"log"
	"net"

	"github.com/andredubov/auth/internal/config"
	"github.com/andredubov/auth/internal/config/env"
	"github.com/andredubov/auth/internal/repository/postgres"
	server "github.com/andredubov/auth/internal/transport/grpc/auth/v1"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/andredubov/auth/pkg/database"
	"github.com/andredubov/auth/pkg/hasher"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	authConfig, err := env.NewAuthConfig()
	if err != nil {
		log.Fatalf("failed to get auth config: %v", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	postgresConfig, err := env.NewPostgresConfig()
	if err != nil {
		log.Fatalf("failed to get postgres config: %v", err)
	}

	connection, err := database.NewPostgresConnection(postgresConfig)
	if err != nil {
		log.Fatalf("failed to get postgres connection: %v", err)
	}
	defer connection.Close()

	usersRepository := postgres.NewUsersRepository(connection)
	passwordHasher := hasher.NewSHA256Hasher(authConfig.PasswordSalt())

	listen, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	auth_v1.RegisterAuthServer(s, server.NewAuthServer(usersRepository, passwordHasher))

	log.Printf("server listening at %v", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
