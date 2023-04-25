package db

import (
	"github.com/go-redis/redis"

	"github.com/cyansilver/go-lib/config"
)

func InitRedisClient(cf *config.AppConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cf.CacheHost,
		Password: cf.CachePwd,
		DB:       0,
	})
}
