package main

import (
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

type FindHandler func(string) (Handler, bool)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send         chan Message
	socket       *websocket.Conn
	findHandler  FindHandler
	session      *r.Session
	stopChannels map[int]chan bool
}

func (c *Client) Close() {
	for _, ch := range c.stopChannels {
		ch <- true
	}
	close(c.send)
}

func (c *Client) NewStopChannel(stopKey int) chan bool {
	stop := make(chan bool)
	c.stopChannels[stopKey] = stop
	return stop
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	for msg := range client.send {
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

// NewClient ...
func NewClient(socket *websocket.Conn, findHandler FindHandler, session *r.Session) *Client {
	return &Client{
		send:         make(chan Message),
		socket:       socket,
		findHandler:  findHandler,
		session:      session,
		stopChannels: make(map[int]chan bool),
	}
}
