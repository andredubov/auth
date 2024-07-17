package repository

import (
	"context"
	"time"
)

const (
	UnknownRole = 0
	UserRole    = 1
	AdminRole   = 2
)

type (
	User struct {
		ID           int64
		Name         string
		Email        string
		PasswordHash string
		UserRole     int
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

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
