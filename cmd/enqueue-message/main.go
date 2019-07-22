// package mqttclient is used as
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const message = `
{
	"identification" : "XXXXXXXXXXXXXXXXXXXX",
	"p1_version" : "50",
	"timestamp" : "190722223037S",
	"equipment_id" : "1234567890123456789012345678901234",
	"energy_delivered_tariff1" : 215.769,
	"energy_delivered_tariff2" : 1904.927,
	"energy_returned_tariff1" : 0,
	"energy_returned_tariff2" : 0,
	"electricity_tariff" : "0002",
	"power_delivered" : 7.654,
	"power_returned" : 0,
	"electricity_failures" : 18,
	"electricity_long_failures" : 2,
	"electricity_failure_log" : "(1)(0-0:96.7.19)(1234567890123)(1234567890*s)",
	"electricity_sags_l1" : 12,
	"electricity_swells_l1" : 2,
	"message_long" : "",
	"voltage_l1" : 234,
	"current_l1" : 3,
	"power_delivered_l1" : 1.654,
	"power_returned_l1" : 0,
	"gas_device_type" : 3,
	"gas_equipment_id" : "1234567890123456789012345678901234",
	"gas_delivered" : 741.87
  }
`

func connect(clientID string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientID, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	return opts
}

func main() {
	uri, err := url.Parse(os.Getenv("CLOUDMQTT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "test"
	}

	fmt.Printf("publishing message to %s/%s\n", uri.Host, string(topic))

	client := connect("pub", uri)
	token := client.Publish(topic, 0, false, message)
	token.Wait()
}
