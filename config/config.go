package config

import (
	"fmt"
	"os"
	"strconv"
)

type dbConfig struct {
	Host string
	User string
	Pass string
	Port int
}

var Redis *dbConfig

func init() {
	Redis = new(dbConfig)
	Redis.Host = os.Getenv("REDIS_HOST")
	Redis.User = os.Getenv("REDIS_USER")
	Redis.Pass = os.Getenv("REDIS_PASS")
	Redis.Port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
}

func (c *dbConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
