package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/auth/internal/client/database"
	"github.com/andredubov/auth/internal/service/model"
)

// Create new user in the repository
func (u *usersRepository) Create(ctx context.Context, user model.User) (int64, error) {
	insertBuilder := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns(nameUsersTableColumn, emailUsersTableColumn, passhashUsersTableColumn, roleUsersTableColumn).
		Values(user.Name, user.Email, user.Password, user.UserRole).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	q := database.Query{
		Name:     "usersRepository.Create",
		QueryRaw: query,
	}

	var userID int64

	err = u.dbClient.Database().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
