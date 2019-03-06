package main

import (
	"fmt"
	"time"

	r "github.com/dancannon/gorethink"
)

func subscribe(session *r.Session, stop <-chan bool) {
	result := make(chan r.ChangeResponse)

	cursor, _ := r.Table("channel").
		Changes().
		Run(session)

	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			result <- change
		}
		fmt.Println("exiting cursor goroutine")
	}()

	for {
		select {
		case change := <-result:
			fmt.Printf("%#v\n", change.NewValue)
		case <-stop:
			fmt.Println("closing cursor")
			cursor.Close()
			return
		}
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
	fmt.Println("sending stop")
	stop <- true
	fmt.Println("browser closes... websocket closes")
	time.Sleep(time.Second * 10000)
}
