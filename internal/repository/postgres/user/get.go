package postgres

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/auth/internal/model"
	"github.com/andredubov/auth/internal/repository"
	"github.com/andredubov/auth/internal/repository/converter"
	modelRepo "github.com/andredubov/auth/internal/repository/model"
	"github.com/jackc/pgx/v4"
)

// Get a user by its email from the repository
func (u *usersRepository) GetByID(ctx context.Context, userID int64) (model.User, error) {
	queryBuilder := sq.Select(
		idUsersTabelColumn,
		nameUsersTableColumn,
		emailUsersTableColumn,
		passhashUsersTableColumn,
		roleUsersTableColumn,
		createdAtUsersTableColumn,
		updatedAtUsersTableColumn,
	).From(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idUsersTabelColumn: userID}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return model.User{}, err
	}

	row, user := u.pool.QueryRow(ctx, query, args...), modelRepo.User{}

	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.UserRole,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, repository.ErrUserNotFound
		}

		return model.User{}, err
	}

	return converter.ToUserFromRepo(&user), nil
}
