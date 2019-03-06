package main

import (
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
)

func subscribe(session *r.Session, stop <-chan bool) {
	var change r.ChangeResponse
	cursor, _ := r.Table("channel").
		Changes().
		Run(session)

	for cursor.Next(&change) {
		fmt.Printf("%#v\n", change.NewValue)
	}
}

func main() {
	session, _ := r.Connect(r.ConnectOpts{
		Address:  "rethinkdb:28015",
		Database: "rtsupport",
	})
	stop := make(chan bool)
	go subscribe(session, stop)
	time.Sleep(time.Second * 5)
	fmt.Println("browser closes... websocket closes")
	time.Sleep(time.Second * 10000)
}
