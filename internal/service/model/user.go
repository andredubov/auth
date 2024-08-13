package model

import (
	"database/sql"
	"time"
)

type Role int32

// User service layer user model
type User struct {
	ID              int64
	Name            string
	Email           string
	Role            Role
	Password        string
	PasswordConfirm string
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

// UpdateUserInfo service layer user model
type UpdateUserInfo struct {
	ID    int64
	Name  *string
	Email *string
	Role  *Role
}
