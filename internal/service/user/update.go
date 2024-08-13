package user

import (
	"context"
	"log"

	"github.com/andredubov/auth/internal/service/model"
)

// Update a user information
func (u *usersService) Update(ctx context.Context, updateUserInfo model.UpdateUserInfo) (int64, error) {
	const op = "usersService.Update:"
	var id int64
	err := u.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = u.usersRepository.Update(ctx, updateUserInfo)
		if errTx != nil {
			log.Printf("%s: %s", op, errTx)
			return errTx
		}

		errTx = u.usersCache.Delete(ctx, updateUserInfo.ID)
		if errTx != nil {
			log.Printf("%s: %s", op, errTx)
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
