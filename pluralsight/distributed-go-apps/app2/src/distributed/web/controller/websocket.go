package controller

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"sync"

	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/dto"
	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/qutils"
	"github.com/GOCODE/pluralsight/distributed-go-apps/app2/src/distributed/web/model"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
)

const url = "amqp://localhost:5672"

type websocketController struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	sockets  []*websocket.Conn
	mutex    sync.Mutex
	upgrader websocket.Upgrader
}

func newWebsocketController() *websocketController {
	wsc := new(websocketController)

	wsc.conn, wsc.ch = qutils.GetChannel(url)

	wsc.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	go wsc.listenForSources()
	go wsc.listenForMessages()

	return wsc
}

func (wsc *websocketController) handleMessage(w http.ResponseWriter, r *http.Request) {
	socket, _ := wsc.upgrader.Upgrade(w, r, nil)
	wsc.addSocket(socket)
	go wsc.listenForDiscoveryRequests(socket)
}

func (wsc *websocketController) addSocket(socket *websocket.Conn) {
	wsc.mutex.Lock()
	wsc.sockets = append(wsc.sockets, socket)
	wsc.mutex.Unlock()
}

func (wsc *websocketController) removeSocket(socket *websocket.Conn) {
	wsc.mutex.Lock()
	socket.Close()

	for i := range wsc.sockets {
		if wsc.sockets[i] == socket {
			wsc.sockets = append(wsc.sockets[:i], wsc.sockets[i+1:]...)
		}
	}

	wsc.mutex.Unlock()
}

func (wsc *websocketController) sendMessage(msg message) {
	socketsToRemove := []*websocket.Conn{}
	for _, socket := range wsc.sockets {
		err := socket.WriteJSON(msg)

		if err != nil {
			socketsToRemove = append(socketsToRemove, socket)
		}
	}

	for _, socket := range socketsToRemove {
		wsc.removeSocket(socket)
	}
}

func (wsc *websocketController) listenForMessages() {
	q := qutils.GetQueue("", wsc.ch, true)
	wsc.ch.QueueBind(
		q.Name, //name string,
		"",     //key string,
		qutils.WebappReadingsExchange, //xchange string,
		false, //noWait bool,
		nil)   //args amqp.Table)

	msgs, _ := wsc.ch.Consume(
		q.Name, //queue string,
		"",     //consumer string,
		true,   //autoAck bool,
		false,  //exclusive bool,
		false,  // noLocal bool,
		false,  // noWait bool,
		nil)    // args amqp.Table)

	for msg := range msgs {
		buf := bytes.NewBuffer(msg.Body)
		dec := gob.NewDecoder(buf)
		sm := dto.SensorMessage{}
		dec.Decode(&sm)

		wsc.sendMessage(message{
			Type: "reading",
			Data: sm,
		})
	}
}

func (wsc *websocketController) listenForSources() {
	q := qutils.GetQueue("", wsc.ch, true)
	wsc.ch.QueueBind(
		q.Name, //name string,
		"",     //key string,
		qutils.WebappSourceExchange, //xchange string,
		false, //noWait bool,
		nil)   //args amqp.Table)

	msgs, _ := wsc.ch.Consume(
		q.Name, //queue string,
		"",     //consumer string,
		true,   //autoAck bool,
		false,  //exclusive bool,
		false,  // noLocal bool,
		false,  // noWait bool,
		nil)    // args amqp.Table)

	for msg := range msgs {
		sensor, _ := model.GetSensorByName(string(msg.Body))
		wsc.sendMessage(message{
			Type: "source",
			Data: sensor,
		})
	}
}

func (wsc *websocketController) listenForDiscoveryRequests(socket *websocket.Conn) {
	for {
		msg := message{}
		err := socket.ReadJSON(&msg)
		if err != nil {
			wsc.removeSocket(socket)
			break
		}

		if msg.Type == "discover" {
			wsc.ch.Publish(
				"", //exchange string,
				qutils.WebappDiscoveryQueue, //key string,
				false,             //mandatory bool,
				false,             //immediate bool,
				amqp.Publishing{}) //msg amqp.Publishing)
		}
	}
}

type message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
