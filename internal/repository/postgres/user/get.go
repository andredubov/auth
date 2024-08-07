package postgres

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/auth/internal/repository"
	"github.com/andredubov/auth/internal/repository/converter"
	modelRepo "github.com/andredubov/auth/internal/repository/model"
	"github.com/andredubov/auth/internal/service/model"
	"github.com/andredubov/golibs/pkg/client/database"
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

	q := database.Query{
		Name:     "usersRepository.GetByID",
		QueryRaw: query,
	}

	user := modelRepo.User{}
	err = u.dbClient.Database().ScanOneContext(ctx, &user, q, args...)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.User{}, repository.ErrUserNotFound
		}
		return model.User{}, err
	}

	return converter.ToUserFromRepo(&user), nil
}
