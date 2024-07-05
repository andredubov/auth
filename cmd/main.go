package main

import (
	"fmt"
	"log"
	"net"

	server "github.com/andredubov/auth/internal/transport/grpc/auth/v1"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

func main() {

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	auth_v1.RegisterAuthServer(s, server.NewAuthServer())

	log.Printf("server listening at %v", listen.Addr())

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
