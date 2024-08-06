package converter

import (
	modelRepo "github.com/andredubov/auth/internal/repository/model"
	"github.com/andredubov/auth/internal/service/model"
)

// ToUserFromRepo converts repository user model to service user model
func ToUserFromRepo(user *modelRepo.User) model.User {
	return model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		UserRole:  user.UserRole,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
