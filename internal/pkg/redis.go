package pkg

/*
--- MIT License (c) 2025 achmad
--- See LICENSE for more details
*/

import (
	"github.com/redis/go-redis/v9"
)

func InitRedisDB(addr string) *redis.Client {
	// init redis db
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return rdb
}
