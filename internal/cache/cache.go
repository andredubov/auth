package cache

import (
	"context"

	"github.com/andredubov/auth/internal/service/model"
)

// Users interface for working with a users cache
type Users interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}
