package models

import (
	"context"
	"github.com/0xc0d/weeconn/db"
	"github.com/0xc0d/weeconn/pkg/suid"
	"github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

func AddUrl(value string) (ID string, err error) {
	var ok bool
	for !ok {
		ID = string(suid.NewUID())
		ok, err = db.NewRedisClient().SetNX(context.Background(), ID, value, 0).Result()
		if err != nil {
			break
		}
	}
	return
}

func GetUrlByID(key string) (value string, err error) {
	value, err = db.NewRedisClient().Get(context.Background(), key).Result()
	return
}
