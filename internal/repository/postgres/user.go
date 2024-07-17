package postgres

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/andredubov/auth/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable = "users"
)

type usersRepository struct {
	pool *pgxpool.Pool
}

// NewUsersRepository create an instance of the usersRepository struct
func NewUsersRepository(pool *pgxpool.Pool) repository.Users {
	return &usersRepository{
		pool,
	}
}

// Create new user in the repository
func (u *usersRepository) Create(ctx context.Context, name, email, password string, role int) (int64, error) {
	insertBuilder := sq.Insert(usersTable).
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "pass_hash", "role").
		Values(name, email, password, role).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	userID := int64(0)

	err = u.pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

// Get a user by its email from the repository
func (u *usersRepository) GetByID(ctx context.Context, userID int64) (repository.User, error) {
	queryBuilder := sq.Select("id", "name", "email", "pass_hash", "role", "created_at", "updated_at").
		From(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userID}).
		Limit(1)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return repository.User{}, err
	}

	row, user := u.pool.QueryRow(ctx, query, args...), repository.User{}

	if err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.UserRole,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return repository.User{}, err
	}

	return user, nil
}

// Update a user in the repository
func (u *usersRepository) Update(ctx context.Context, userInfo repository.UpdateUserInfo) (int64, error) {
	updateBuilder := sq.Update(usersTable).
		PlaceholderFormat(sq.Dollar)

	if userInfo.Name != nil {
		updateBuilder.Set("name", userInfo.Name)
	}

	if userInfo.Email != nil {
		updateBuilder.Set("email", userInfo.Email)
	}

	if userInfo.UserRole != nil {
		updateBuilder.Set("role", userInfo.UserRole)
	}

	updateBuilder.Set("updated_at", time.Now()).Where(sq.Eq{"id": userInfo.ID})

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

// Delete a user from the repository
func (u *usersRepository) Delete(ctx context.Context, userID int64) (int64, error) {
	deleteBuilder := sq.Delete(usersTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userID})

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
