package db

import (
	"github.com/0xc0d/weeconn/config"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

var once sync.Once
var rdb *redis.Client

func NewRedisClient() *redis.Client {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     config.Redis.Addr(),
			Username: config.Redis.User,
			Password: config.Redis.Pass,
		})

		log.Printf("Connecting to Redis at %s", config.Redis.Addr())
	})
	return rdb
}
