package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
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
		msgType, msg, err := socket.ReadMessage()
		logFatal(err)
		fmt.Println(string(msg))
		err = socket.WriteMessage(msgType, msg)
		logFatal(err)
	}

}
