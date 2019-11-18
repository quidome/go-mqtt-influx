// Package mqttagent handles this projects mqtt needs
package mqttagent

import (
	"fmt"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("pkg", "mqttagent")

// Listen initiates the client connection
func Listen(uri *url.URL, topic string, messageChannel chan []byte) {
	opts := createClientOptions("mqttinflux", uri, func(client mqtt.Client) {
		createSubscriptions(client, messageChannel, topic)
	})
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
}

func createSubscriptions(client mqtt.Client, messageChannel chan []byte, topic string) {
	log.Info("creating subscriptions")
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		messageChannel <- msg.Payload()
	})
}

func createClientOptions(clientID string, uri *url.URL, connectHandler func(client mqtt.Client)) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetClientID(clientID)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(time.Second)
	opts.SetOnConnectHandler(connectHandler)
	return opts
}
