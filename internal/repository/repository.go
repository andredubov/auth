package repository

import (
	"context"
	"errors"
	"time"
)

var (
	// error user not found
	ErrUserNotFound = errors.New("User not found")
)

type (
	// Output user data from the user's repository
	User struct {
		ID           int64
		Name         string
		Email        string
		PasswordHash string
		UserRole     int
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	// Input user data to update
	UpdateUserInfo struct {
		ID       int64
		Name     *string
		Email    *string
		UserRole *int
	}
)

// Users users repository interface
type Users interface {
	Create(ctx context.Context, name, email, password string, role int) (int64, error)
	GetByID(ctx context.Context, useID int64) (User, error)
	Update(ctx context.Context, userInfo UpdateUserInfo) (int64, error)
	Delete(ctx context.Context, userID int64) (int64, error)
}
