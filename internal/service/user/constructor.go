package user

import (
	"github.com/andredubov/auth/internal/client/database"
	"github.com/andredubov/auth/internal/repository"
	"github.com/andredubov/auth/internal/service"
	"github.com/andredubov/auth/pkg/hasher"
)

type usersService struct {
	usersRepository repository.Users
	hasher          hasher.PasswordHasher
	txManager       database.TxManager
}

func NewService(
	usersRepository repository.Users,
	hasher hasher.PasswordHasher,
	txManager database.TxManager,
) service.Users {
	return &usersService{
		usersRepository,
		hasher,
		txManager,
	}
}
