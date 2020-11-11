package config

import (
	"fmt"
	"os"
)

var Redis struct {
	Host     string
	Username string
	Password string
}

func init() {
	host := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	Redis.Host = host
	Redis.Username = os.Getenv("REDIS_USER")
	Redis.Password = os.Getenv("REDIS_PASS")
}
