package db

import (
	"github.com/0xc0d/weeconn/config"
	influxdb "github.com/influxdata/influxdb/client/v2"
	"log"
	"sync"
)

var newInflux struct {
	once   sync.Once
	client influxdb.Client
}

func NewInfluxdbClient() influxdb.Client {
	newInflux.once.Do(func() {
		log.Printf("Connecting to Influx at %s", config.Influx.URL)

		var err error
		newInflux.client, err = influxdb.NewHTTPClient(influxdb.HTTPConfig{
			Addr:     config.Influx.URL,
			Username: config.Influx.Username,
			Password: config.Influx.Password,
		})
		if err != nil {
			log.Println(err)
		}

	})
	return newInflux.client
}
