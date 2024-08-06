package model

import (
	"database/sql"
	"time"
)

// User service layer user model
type User struct {
	ID              int64
	Name            string
	Email           string
	UserRole        int
	Password        string
	PasswordConfirm string
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

// UpdateUserInfo service layer user model
type UpdateUserInfo struct {
	ID       int64
	Name     *string
	Email    *string
	UserRole *int
}
