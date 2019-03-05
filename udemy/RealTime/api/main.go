package main

import (
	"fmt"
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
)

type Channel struct {
	ID   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
}

type User struct {
	ID   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "rethinkdb:28015",
		Database: "rtsupport",
	})
	logPanic(err)

	router := NewRouter(session)
	fmt.Println("router created!")

	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)
	http.Handle("/", router)
	http.ListenAndServe(":3001", nil)
}

func logFatal(err error) {
	if err != nil {
		log.Println(err)
	}
}

func logPanic(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}
