package user

import (
	"context"
	"log"

	"github.com/andredubov/auth/internal/service/model"
)

// Create a new user
func (u *usersService) Create(ctx context.Context, user model.User) (int64, error) {
	const op = "usersService.Create:"

	hashedPassword, err := u.hasher.HashAndSalt(user.Password)
	if err != nil {
		log.Printf("%s: %s", op, err)
		return 0, err
	}

	user.Password = hashedPassword

	id, err := u.usersRepository.Create(ctx, user)
	if err != nil {
		log.Printf("%s: %s", op, err)
		return 0, err
	}

	return id, err
}
