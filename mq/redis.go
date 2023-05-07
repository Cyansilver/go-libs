package mq

import (
	"github.com/go-redis/redis"

	"github.com/cyansilver/go-libs/config"
	"github.com/cyansilver/go-libs/db"
)

// Producer wrapper the redis.Client
type Producer = redis.Client

// Consumer wrapper the redis.Client
type Consumer = redis.Client

func InitProducer(cf *config.AppConfig) *Producer {
	return db.InitRedisClient(cf)
}

func InitConsumer(cf *config.AppConfig) *Consumer {
	return db.InitRedisClient(cf)
}
