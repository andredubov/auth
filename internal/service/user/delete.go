package user

import (
	"context"
	"log"
)

func (u *usersService) Delete(ctx context.Context, userID int64) (int64, error) {
	const op = "usersService.Delete:"

	id, err := u.usersRepository.Delete(ctx, userID)
	if err != nil {
		log.Printf("%s: %s", op, err)
		return 0, err
	}

	return id, nil
}
