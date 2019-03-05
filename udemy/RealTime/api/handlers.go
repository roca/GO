package main

import (
	"fmt"

	r "github.com/dancannon/gorethink"
	"github.com/mitchellh/mapstructure"
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
	go func() {
		cursor, err := r.Table("channel").
			Changes(r.ChangesOpts{IncludeInitial: true}).
			Run(client.session)
		if err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}
		var change r.ChangeResponse
		for cursor.Next(&change) {
			if change.NewValue != nil && change.OldValue == nil {
				client.send <- Message{"channel add", change.NewValue}
				fmt.Println("sent channel add msg")
			}
		}
	}()

}
