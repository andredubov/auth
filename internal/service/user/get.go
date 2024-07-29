package user

import (
	"context"
	"log"

	"github.com/andredubov/auth/internal/model"
)

func (u *usersService) GetByID(ctx context.Context, userID int64) (model.User, error) {
	const op = "usersService.GetByID:"

	user, err := u.usersRepository.GetByID(ctx, userID)
	if err != nil {
		log.Printf("%s: %s", op, err)
		return model.User{}, err
	}

	return user, nil
}
