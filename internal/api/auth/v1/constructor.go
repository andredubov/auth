package server

import (
	"github.com/andredubov/auth/internal/service"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
)

type Implementation struct {
	auth_v1.UnimplementedAuthServer
	usersService service.Users
}

func NewImplementation(service service.Users) *Implementation {
	return &Implementation{
		usersService: service,
	}
}
