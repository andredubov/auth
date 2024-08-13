package user

import (
	"context"
	"log"
)

// Delete is used to delete a user by its id
func (u *usersService) Delete(ctx context.Context, userID int64) (int64, error) {
	const op = "usersService.Delete:"
	var id int64
	err := u.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = u.usersRepository.Delete(ctx, userID)
		if errTx != nil {
			log.Printf("%s: %s", op, errTx)
			return errTx
		}

		errTx = u.usersCache.Delete(ctx, userID)
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
