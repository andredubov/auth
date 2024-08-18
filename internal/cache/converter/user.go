package converter

import (
	"database/sql"
	"time"

	modelCache "github.com/andredubov/auth/internal/cache/model"
	"github.com/andredubov/auth/internal/service/model"
)

func ToUserCacheFromModel(user model.User) *modelCache.User {
	var updatedAtNs *int64
	if user.UpdatedAt.Valid {
		updatedAtNs = new(int64)
		*updatedAtNs = user.UpdatedAt.Time.Unix()
	}

	return &modelCache.User{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        modelCache.Role(user.Role),
		CreatedAtNs: user.UpdatedAt.Time.Unix(),
		UpdatedAtNs: updatedAtNs,
	}
}

func ToUserFromCache(user *modelCache.User) *model.User {
	var updatedAt sql.NullTime

	if user.UpdatedAtNs != nil {
		updatedAt = sql.NullTime{
			Time:  time.Unix(0, *user.UpdatedAtNs),
			Valid: true,
		}
	}
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      model.Role(user.Role),
		CreatedAt: time.Unix(0, user.CreatedAtNs),
		UpdatedAt: updatedAt,
	}
}
