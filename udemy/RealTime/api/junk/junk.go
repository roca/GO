package main

import (
	"fmt"
	"log"

	r "github.com/dancannon/gorethink"
)

type User struct {
	ID   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "rethinkdb:28015",
		Database: "rtsupport",
	})
	logError(err)
	// user := User{Name: "anonymous"}
	// response, err := r.Table("user").
	// 	Insert(user).
	// 	RunWrite(session)
	// logError(err)

	//user := User{Name: "Jame Moore"}
	// response, _ := r.Table("user").
	// 	Delete().
	// 	RunWrite(session)
	// fmt.Printf("%#v\n", response)

	cursor, _ := r.Table("user").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(session)
	var changeResponse r.ChangeResponse
	for cursor.Next(&changeResponse) {
		fmt.Printf("%#v\n", changeResponse)
	}

}
