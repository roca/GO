package main

import (
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
