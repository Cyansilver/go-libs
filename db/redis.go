package db

import (
	"github.com/cyansilver/go-libs/config"
	"github.com/go-redis/redis"
)

func InitRedisClient(cf *config.AppConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cf.CacheHost,
		Password: cf.CachePwd,
		DB:       0,
	})
}
