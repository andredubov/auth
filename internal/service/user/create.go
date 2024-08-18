package user

import (
	"context"
	"log"

	"github.com/andredubov/auth/internal/service/model"
)

// Create a new user
func (u *usersService) Create(ctx context.Context, user model.User) (int64, error) {
	const op = "usersService.Create:"
	var id int64
	err := u.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		hashedPassword, errTx := u.hasher.HashAndSalt(user.Password)
		if errTx != nil {
			log.Printf("%s: %s", op, errTx)
			return errTx
		}

		user.Password = hashedPassword

		id, errTx = u.usersRepository.Create(ctx, user)
		if errTx != nil {
			log.Printf("%s: %s", op, errTx)
			return errTx
		}

		user.ID = id

		errTx = u.usersCache.Create(ctx, user)
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
