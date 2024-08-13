package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/golibs/pkg/client/database"
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

	q := database.Query{
		Name:     "usersRepository.Delete",
		QueryRaw: query,
	}

	res, err := u.dbClient.Database().ExecContext(ctx, q, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
