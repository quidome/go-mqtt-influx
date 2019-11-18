package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/quidome/mqtt-influx/config"
	"github.com/quidome/mqtt-influx/dsmr"
	"github.com/quidome/mqtt-influx/influxagent"
	"github.com/quidome/mqtt-influx/mqttagent"
	"github.com/sirupsen/logrus"
)

const banner = ` _______  _______ __________________       _________ _        _______  _
(       )(  ___  )\__   __/\__   __/       \__   __/( (    /|(  ____ \( \      |\     /||\     /|
| () () || (   ) |   ) (      ) (             ) (   |  \  ( || (    \/| (      | )   ( |( \   / )
| || || || |   | |   | |      | |    _____    | |   |   \ | || (__    | |      | |   | | \ (_) /
| |(_)| || |   | |   | |      | |   (_____)   | |   | (\ \) ||  __)   | |      | |   | |  ) _ (
| |   | || | /\| |   | |      | |             | |   | | \   || (      | |      | |   | | / ( ) \
| )   ( || (_\ \ |   | |      | |          ___) (___| )  \  || )      | (____/\| (___) |( /   \ )
|/     \|(____\/_)   )_(      )_(          \_______/|/    )_)|/       (_______/(_______)|/     \|
in to the flux
`

var log = logrus.WithField("pkg", "main")

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true})

	fmt.Print(banner)

	// load settings (from env vars)
	settings := config.Get()

	fmt.Println("configured to:")
	fmt.Printf("read messages from: %s/%s\n", settings.MQTTURL, settings.MQTTTopic)
	fmt.Printf("and send them to:   %s/%s\n\n", settings.InfluxURL, settings.InfluxDB)

	mqttURI, err := url.Parse(settings.MQTTURL)
	if err != nil {
		log.Fatal(err)
	}

	// create channel for dsmr telegrams
	dsmrChannel := make(chan []byte)

	// create mqtt listener
	fmt.Printf("Connect to %s/%s\n", mqttURI, settings.MQTTTopic)
	mqttagent.Listen(mqttURI, settings.MQTTTopic, dsmrChannel)

	// get messages and send them through
	for {
		// get message from listener
		message := <-dsmrChannel

		//  transform into data structure
		res := dsmr.Telegram{}
		err := json.Unmarshal(message, &res)
		if err != nil {
			log.Error(err)
		}

		// send data to influx
		influxagent.StoreData(settings.InfluxURL, settings.InfluxDB, settings.InfluxUser, settings.InfluxPass, res)
	}
}
