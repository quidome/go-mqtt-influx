package config

import (
	"github.com/spf13/viper"
)

var (
	mngr = viper.New()
	//log  = logger.New()
)

// Settings contains application settings
type Settings struct {
	MQTTURL    string
	MQTTTopic  string
	InfluxURL  string
	InfluxDB   string
	InfluxUser string
	InfluxPass string
}

const (
	mqttURL    = "MQTT_URL"
	mqttTopic  = "MQTT_TOPIC"
	influxURL  = "INFLUX_URL"
	influxDB   = "INFLUX_DB"
	influxUser = "INFLUX_USER"
	influxPass = "INFLUX_PASS"
)

func init() {
	mngr.SetDefault(mqttURL, "tcp://localhost:1883")
	mngr.SetDefault(mqttTopic, "test")
	mngr.SetDefault(influxURL, "http://localhost:8086")
	mngr.SetDefault(influxDB, "testdb")
	mngr.SetDefault(influxUser, "testuser")
	mngr.SetDefault(influxPass, "supersecret")

	mngr.AutomaticEnv()
}

// Get the application settings from various source configs
func Get() Settings {
	return Settings{
		MQTTURL:    mngr.GetString(mqttURL),
		MQTTTopic:  mngr.GetString(mqttTopic),
		InfluxURL:  mngr.GetString(influxURL),
		InfluxDB:   mngr.GetString(influxDB),
		InfluxUser: mngr.GetString(influxUser),
		InfluxPass: mngr.GetString(influxPass),
	}
}
