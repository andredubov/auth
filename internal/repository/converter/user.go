package converter

import (
	"github.com/andredubov/auth/internal/model"
	modelRepo "github.com/andredubov/auth/internal/repository/model"
)

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
