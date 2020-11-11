package db

import (
	"github.com/0xc0d/weeconn/config"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

var newRedis struct {
	once   sync.Once
	client *redis.Client
}

func NewRedisClient() *redis.Client {
	newRedis.once.Do(func() {
		newRedis.client = redis.NewClient(&redis.Options{
			Addr:     config.Redis.Host,
			Username: config.Redis.Username,
			Password: config.Redis.Password,
		})

		log.Printf("Connecting to Redis at %s", config.Redis.Host)
	})
	return newRedis.client
}
