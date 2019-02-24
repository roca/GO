package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3001", nil)
}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello from GO")
	socket, err := upgrader.Upgrade(w, r, nil)
	logFatal(err)
	for {
		// msgType, msg, err := socket.ReadMessage()
		// logFatal(err)
		var inMessage Message
		var outMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			logFatal(err)
			break
		}
		switch inMessage.Name {
		case "channel add":
			if err := addChannel(inMessage.Data); err != nil {
				outMessage = Message{"error", err}
				if err := socket.WriteJSON(&outMessage); err != nil {
					logFatal(err)
					break
				}
			}
		}
		// fmt.Println(string(msg))
		// err = socket.WriteMessage(msgType, msg)
		// logFatal(err)
	}

}

func addChannel(data interface{}) error {
	var channel Channel

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.ID = "1"
	fmt.Printf("%#v\n", channel)
	return nil
}
