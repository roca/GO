package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
	"github.com/mitchellh/mapstructure"
)

const (
	ChannelStop = iota
	UserStop
	MessageStop
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	if err := mapstructure.Decode(data, &channel); err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}

	go func() {
		if err := r.Table("channel").
			Insert(channel).
			Exec(client.session); err != nil {
			client.send <- Message{"error", err.Error()}
		}
	}()
}

func subscribeChannel(client *Client, data interface{}) {
	stop := client.NewStopChannel(ChannelStop)
	result := make(chan r.ChangeResponse)

	cursor, err := r.Table("channel").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(client.session)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}

	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			result <- change
		}
	}()

	go func() {
		for {
			select {
			case <-stop:
				cursor.Close()
				return
			case change := <-result:
				if change.NewValue != nil && change.OldValue == nil {
					client.send <- Message{"channel add", change.NewValue}
					fmt.Println("sent channel add msg")
				}
			}
		}
	}()

}

func unsubscribeChannel(client *Client, data interface{}) {
	client.StopForKey(ChannelStop)
}

// User Events
func editUser(client *Client, data interface{}) {}

func subscribeUser(client *Client, data interface{}) {}

func unsubscribeUser(client *Client, data interface{}) {
	client.StopForKey(UserStop)
}

func addUser(client *Client, data interface{}) {}

func removeUser(client *Client, data interface{}) {}

//Message Events
func addMessage(client *Client, data interface{}) {}

func subscribeMessage(client *Client, data interface{}) {}

func unsubscribeMessage(client *Client, data interface{}) {
	client.StopForKey(MessageStop)
}
