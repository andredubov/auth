package postgres

import (
	"github.com/andredubov/auth/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable                = "users"
	idUsersTabelColumn        = "id"
	nameUsersTableColumn      = "name"
	emailUsersTableColumn     = "email"
	passhashUsersTableColumn  = "pass_hash"
	roleUsersTableColumn      = "role"
	createdAtUsersTableColumn = "created_at"
	updatedAtUsersTableColumn = "updated_at"
	unknownRole               = 0
	userRole                  = 1
	adminRole                 = 2
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
