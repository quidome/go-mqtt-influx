package main

import (
	"fmt"

	"github.com/quidome/mqtt-influx/config"
)

const banner = `
 _______  _______ __________________       _________ _        _______  _
(       )(  ___  )\__   __/\__   __/       \__   __/( (    /|(  ____ \( \      |\     /||\     /|
| () () || (   ) |   ) (      ) (             ) (   |  \  ( || (    \/| (      | )   ( |( \   / )
| || || || |   | |   | |      | |    _____    | |   |   \ | || (__    | |      | |   | | \ (_) /
| |(_)| || |   | |   | |      | |   (_____)   | |   | (\ \) ||  __)   | |      | |   | |  ) _ (
| |   | || | /\| |   | |      | |             | |   | | \   || (      | |      | |   | | / ( ) \
| )   ( || (_\ \ |   | |      | |          ___) (___| )  \  || )      | (____/\| (___) |( /   \ )
|/     \|(____\/_)   )_(      )_(          \_______/|/    )_)|/       (_______/(_______)|/     \|
fl to the ux
`

func main() {
	fmt.Print(banner)

	// do config stuff
	settings := config.Get()

	fmt.Print(settings)

	// // get mqtt server from
	// uri, err := url.Parse(os.Getenv("CLOUDMQTT_URL"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// a := &telegram.P1Message{}

	// fmt.Print(*a)
	// // connect to mqtt server
}
