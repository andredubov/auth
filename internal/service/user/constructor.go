package user

import (
	"github.com/andredubov/auth/internal/repository"
	"github.com/andredubov/auth/internal/service"
	"github.com/andredubov/golibs/pkg/client/database"
	"github.com/andredubov/golibs/pkg/hasher"
)

type usersService struct {
	usersRepository repository.Users
	hasher          hasher.PasswordHasher
	txManager       database.TxManager
}

// NewService creates a instance of usersService struct
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
