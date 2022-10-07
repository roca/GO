package main

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/nats-io/nats.go"
)

var favorites = map[string]string{
	"color":  "blue",
	"food":   "peanut butter",
	"season": "summer",
	"movie":  "The Guns of Navarone",
	"album":  "Giant Steps",
}

func main() {
	natsUrl := os.Getenv("NATS_URL")
	if natsUrl == "" {
		natsUrl = "nats://demo.nats.io:4222"
	}

	nc, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatalln(err)
	}
	//defer nc.Close()
	nc.QueueSubscribe("gophercon.services", "romel", func(msg *nats.Msg) {
		msg.Respond([]byte("gophercon.services.romel.favorites"))
	})
	nc.QueueSubscribe("gophercon.services", "romel", func(msg *nats.Msg) {
		val, ok := favorites[string(msg.Data)]
		if ok {
			keys := []string{}
			for k := range favorites {
				keys = append(keys, k)
			}
			msg.Respond(
				[]byte("Romels favorites" + strings.Join(keys, ",")),
			)
			return
		}
		msg.Respond([]byte(val))
	})

	runtime.Goexit()
}
