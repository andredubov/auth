package model

import (
	"database/sql"
	"time"
)

// User repository layer user model
type User struct {
	ID           int64
	Name         string
	Email        string
	UserRole     int
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
}

// UpdateUserInfo repository layer user update info model
type UpdateUserInfo struct {
	ID       int64
	Name     *string
	Email    *string
	UserRole *int
}
