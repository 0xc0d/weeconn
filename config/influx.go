package config

import (
	"fmt"
	"os"
)

var Influx struct {
	Database   string
	URL      string
	Username string
	Password string
}

func init() {
	url := fmt.Sprintf("http://%s:%s", os.Getenv("INFLUX_HOST"), os.Getenv("INFLUX_PORT"))
	Influx.URL = url
	Influx.Username = os.Getenv("INFLUX_USER")
	Influx.Password = os.Getenv("INFLUX_PASS")
	Influx.Database = os.Getenv("INFLUX_DB")
}
