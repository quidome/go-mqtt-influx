package main

import (
	"fmt"
	"github.com/influxdb/influxdb/client"
	"log"
	"net/url"
)

var ic *client.Client

func connect(host string, port string, username string, password string) {

	u, err := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}

	ic, err = client.NewClient(client.Config{URL: *u})
	if err != nil {
		log.Fatal(err)
	}

	if _, _, err := ic.Ping(); err != nil {
		log.Fatal(err)
	}

	ic.SetAuth(username, password)
}

func main() {
	connect("localhost", "8086", "testuser", "pass")
}
