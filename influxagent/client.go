package influxagent

import (
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/quidome/mqtt-influx/dsmr"
)

// StoreData takes a bunch of parameters to construct and send a message to influxdb
func StoreData(influxURL string, influxDatabase string, influxUsername string, influxPassword string, influxData dsmr.Telegram) {
	// A dsmr.Telegram is received. The values in the telegram need to be send to influx
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     influxURL,
		Username: influxUsername,
		Password: influxPassword,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	} else {
		fmt.Printf("Sending metrics to %s/%s as %s\n", influxURL, influxDatabase, influxUsername)
	}
	defer c.Close()

	// create new batch point
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influxDatabase,
		Precision: "s",
	})

	// Create a point and add to batch
	point_tags := map[string]string{"location": "ground"}

	pt, err := client.NewPoint("p1_meter", point_tags, structs.Map(influxData), time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)
	c.Write(bp)
}
