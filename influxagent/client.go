package influxagent

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/fatih/structs"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/quidome/mqtt-influx/dsmr"
)

var log = logrus.WithField("pkg", "influxagent")

// StoreData takes a bunch of parameters to construct and send a message to influxdb
func StoreData(influxURL string, influxDatabase string, influxUsername string, influxPassword string, influxData dsmr.Telegram) {
	// A dsmr.Telegram is received. The values in the telegram need to be send to influx
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     influxURL,
		Username: influxUsername,
		Password: influxPassword,
	})
	if err != nil {
		log.Errorf("Error creating InfluxDB Client: %s", err.Error())
		return
	}
	defer c.Close()

	// create new batch point
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influxDatabase,
		Precision: "s",
	})

	// Create a point and add to batch
	pointTags := map[string]string{"location": "ground"}

	pt, err := client.NewPoint("p1_meter", pointTags, structs.Map(influxData), time.Now())
	if err != nil {
		log.Errorf("Error creating new point: %s", err.Error())
		return
	}
	bp.AddPoint(pt)
	if err = c.Write(bp); err != nil {
		log.Errorf("Error writing data %s", err)
		return
	}
}
