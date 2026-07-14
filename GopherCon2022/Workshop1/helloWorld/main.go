package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalln(err)
	}
	nc.Subscribe("hello.world", func(m *nats.Msg) {
		log.Printf("Received a message on : %s \n", m.Subject)
		log.Printf("Received a message: %s \n", string(m.Data))
		m.Respond([]byte("Hey dude"))
	})
	fmt.Println("Subscribed to " + string(nc.ConnectedUrl()))

	runtime.Goexit()
}
