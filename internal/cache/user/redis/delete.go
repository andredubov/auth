package redis

import (
	"context"
	"strconv"
)

func (u *usersCache) Delete(ctx context.Context, id int64) error {
	idStr := strconv.FormatInt(id, 10)
	err := u.cacheClient.Cache().Delete(ctx, idStr)
	if err != nil {
		return err
	}

	return nil
}
