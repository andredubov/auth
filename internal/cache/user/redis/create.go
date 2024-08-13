package redis

import (
	"context"
	"strconv"

	"github.com/andredubov/auth/internal/cache/converter"
	"github.com/andredubov/auth/internal/service/model"
)

func (u *usersCache) Create(ctx context.Context, user *model.User) error {
	userCache := converter.ToUserCacheFromModel(user)
	hash := strconv.FormatInt(userCache.ID, 10)

	err := u.cacheClient.Cache().HashSet(ctx, hash, userCache)
	if err != nil {
		return err
	}

	return nil
}
