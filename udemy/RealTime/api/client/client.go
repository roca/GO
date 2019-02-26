package main

import "fmt"

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send chan Message
}

func (client *Client) write() {
	for msg := range client.send {
		fmt.Printf("%#v\n", msg)
	}
}

func main() {

}
