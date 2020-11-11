package models

import (
	"encoding/json"
	"fmt"
	"github.com/0xc0d/weeconn/config"
	"github.com/0xc0d/weeconn/db"
	influxdb "github.com/influxdata/influxdb/client/v2"
)

// Seen increments an ID view by one
func Seen(id string) {
	batchPoints, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database: config.Influx.Database,
	})
	if err != nil {
		return
	}
	tag := map[string]string{"id": id}
	field := map[string]interface{}{"seen": true}
	point, err := influxdb.NewPoint("seen", tag, field)
	if err != nil {
		return
	}
	batchPoints.AddPoint(point)

	db.NewInfluxdbClient().Write(batchPoints)
}

// Stat returns how many times an ID has been viewed in pass [dayPeriod]days
func Stat(id string, dayPeriod int) (seen int64, err error) {
	query := influxdb.Query{
		Command: fmt.Sprintf("SELECT COUNT(seen) FROM seen WHERE id = '%s'", id),
		Database: config.Influx.Database,
	}
	if dayPeriod != 0 {
		query.Command += fmt.Sprintf("AND time > now() - %dd", dayPeriod)
	}
	response, err := db.NewInfluxdbClient().Query(query)
	if err != nil {
		return
	}

	if len(response.Results[0].Series) == 0 {
		return
	}

	seen, err = response.Results[0].Series[0].Values[0][1].(json.Number).Int64()
	return
}
