package model

import (
	"database/sql"
	"time"
)

type Role int32

// User repository layer user model
type User struct {
	ID           int64        `db:"id"`
	Name         string       `db:"name"`
	Email        string       `db:"email"`
	Role         Role         `db:"role"`
	PasswordHash string       `db:"pass_hash"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at"`
}

// UpdateUserInfo repository layer user update info model
type UpdateUserInfo struct {
	ID    int64
	Name  *string
	Email *string
	Role  *Role
}
