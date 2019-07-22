# go-mqtt-influx
Reads json messages from MQTT queue and writes to influx.

## use local docker

```sh
$ docker-compose up -d
```
## push message to queue

```sh
$  CLOUDMQTT_URL=mqtt://localhost:1883/ go run cmd/enqueue-message/main.go
```