package user

import (
	"context"
	"log"

	"github.com/andredubov/auth/internal/service/model"
)

// Update a user information
func (u *usersService) Update(ctx context.Context, updateUserInfo model.UpdateUserInfo) (int64, error) {
	const op = "usersService.Update:"

	id, err := u.usersRepository.Update(ctx, updateUserInfo)
	if err != nil {
		log.Printf("%s: %s", op, err)
		return 0, err
	}

	return id, nil
}
