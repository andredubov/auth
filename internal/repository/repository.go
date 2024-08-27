package repository

import (
	"context"
	"errors"

	"github.com/andredubov/auth/internal/service/model"
)

var (
	// ErrUserNotFound is an error user not found
	ErrUserNotFound = errors.New("user not found")
)

// Users interface for working with a users repository
type Users interface {
	Create(ctx context.Context, user model.User) (int64, error)
	GetByID(ctx context.Context, userID int64) (*model.User, error)
	Update(ctx context.Context, updateUserInfo model.UpdateUserInfo) (int64, error)
	Delete(ctx context.Context, userID int64) (int64, error)
}
