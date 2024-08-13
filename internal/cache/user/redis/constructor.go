package redis

import (
	"github.com/andredubov/auth/internal/cache"
	cacheClient "github.com/andredubov/golibs/pkg/client/cache"
)

type usersCache struct {
	cacheClient cacheClient.Client
}

func NewUsersCache(cacheClient cacheClient.Client) cache.Users {
	return &usersCache{
		cacheClient,
	}
}
