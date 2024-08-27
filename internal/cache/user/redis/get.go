package redis

import (
	"context"
	"errors"
	"strconv"

	"github.com/andredubov/auth/internal/cache/converter"
	modelCache "github.com/andredubov/auth/internal/cache/model"
	"github.com/andredubov/auth/internal/service/model"
	redigo "github.com/gomodule/redigo/redis"
)

func (u *usersCache) Get(ctx context.Context, id int64) (*model.User, error) {
	idStr := strconv.FormatInt(id, 10)
	values, err := u.cacheClient.Cache().HashGetAll(ctx, idStr)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, errors.New("user not found")
	}

	var user modelCache.User
	if err = redigo.ScanStruct(values, &user); err != nil {
		return nil, err
	}

	return converter.ToUserFromCache(&user), nil
}
