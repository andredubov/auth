package user

import (
	"context"
	"log"

	"github.com/andredubov/auth/internal/service/model"
)

// GetByID gets a user by its id
func (u *usersService) GetByID(ctx context.Context, userID int64) (*model.User, error) {
	const op = "usersService.GetByID:"
	var user *model.User
	err := u.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		user, errTx = u.usersCache.Get(ctx, userID)
		if errTx != nil {
			user, errTx = u.usersRepository.GetByID(ctx, userID)
			if errTx != nil {
				log.Printf("%s: %s", op, errTx)
				return errTx
			}

			errTx = u.usersCache.Create(ctx, *user)
			if errTx != nil {
				log.Printf("%s: %s", op, errTx)
				return errTx
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
