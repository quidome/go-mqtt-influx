package main

import (
	"fmt"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

func main() {
	myHTTPInfluxAddress := "http://127.0.0.1:8086"

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: myHTTPInfluxAddress,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	// create new batch point
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "testdb",
		Precision: "s",
	})

	// Create a point and add to batch
	//point_tags := {}
	point_tags := map[string]string{"location": "ground"}
	point_fields := map[string]interface{}{
		"temp_c":   6,
		"humidity": 23,
	}
	pt, err := client.NewPoint("p1_meter", point_tags, point_fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)
	c.Write(bp)
}
