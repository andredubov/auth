package postgres

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/auth/internal/service/model"
)

// Update a user in the repository
func (u *usersRepository) Update(ctx context.Context, userInfo model.UpdateUserInfo) (int64, error) {
	updateBuilder := sq.Update(usersTable).
		PlaceholderFormat(sq.Dollar)

	if userInfo.Name != nil {
		updateBuilder = updateBuilder.Set(nameUsersTableColumn, userInfo.Name)
	}

	if userInfo.Email != nil {
		updateBuilder = updateBuilder.Set(emailUsersTableColumn, userInfo.Email)
	}

	if userInfo.UserRole != nil {
		updateBuilder = updateBuilder.Set(roleUsersTableColumn, userInfo.UserRole)
	}

	updateBuilder = updateBuilder.Set(updatedAtUsersTableColumn, time.Now()).Where(sq.Eq{idUsersTabelColumn: userInfo.ID})

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	res, err := u.pool.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
