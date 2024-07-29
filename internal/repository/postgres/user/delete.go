package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// Delete a user from the repository
func (u *usersRepository) Delete(ctx context.Context, userID int64) (int64, error) {
	deleteBuilder := sq.Delete(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idUsersTabelColumn: userID})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	res, err := u.pool.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
