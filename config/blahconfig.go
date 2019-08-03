package config

import (
	"github.com/spf13/viper"
)

// Settings contains application settings
type Settings struct {
	MQTTURL string

	// InfluxDBHost string
	// InfluxDBUser string
	// InfluxDBPass string
}

var (
	mngr = viper.New()
)

const (
	mqttURL = "MQTT_URL"
	// gcpProjectID = "GCP_PROJECT_ID"
)

func init() {
	mngr.SetDefault(mqttURL, "mqtt://localhost:1883")

	// _, mqttURL := os.LookupEnv("MQTT_URL")
	// if mqttURL {
	// 	fmt.Println("got from environment: ", mqttURL)
	// }

	// mngr.SetConfigFile("./config/config.yml")
	// err = mngr.ReadInConfig()
	// if err != nil {
	// 	log.Warn("unable to read config file", zap.Error(err))
	// }
}

// Get the application settings from various source configs
func Get() Settings {
	return Settings{
		MQTTURL: mngr.GetString(mqttURL),
		// IncomingMessageSubscription: mngr.GetString(subscription),
		// GCPProjectID:                mngr.GetString(gcpProjectID),
	}
}
