package postgres

import (
	"github.com/andredubov/auth/internal/client/database"
	"github.com/andredubov/auth/internal/repository"
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
	dbClient database.Client
}

// NewUsersRepository create an instance of the usersRepository struct
func NewUsersRepository(dbClient database.Client) repository.Users {
	return &usersRepository{
		dbClient,
	}
}
