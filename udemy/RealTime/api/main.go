package main

import (
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

/*

Channel Events
    'channel add'
    'channel subscribe'
    'channel unsubscribe'

User events
    'user edit'
    'user subscribe'
    'user unsubscribe'
    'user add'
    'user remove'

Message Events
    'message add'
    'message subscribe'
    'message unsubscribe'

*/

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "rethinkdb:28015",
		Database: "rtsupport",
	})
	logPanic(err)

	router := NewRouter(session)

	// Channel Events
	router.Handle("channel add", addChannel)
	router.Handle("channel subscribe", subscribeChannel)
	router.Handle("channel unsubscribe", unsubscribeChannel)

	// User events
	router.Handle("user edit", editUser)
	router.Handle("user subscribe", subscribeUser)
	router.Handle("user unsubscribe", unsubscribeUser)
	router.Handle("user add", addUser)
	router.Handle("user remove", removeUser)

	//Message Events
	router.Handle("message add", addMessage)
	router.Handle("message subscribe", subscribeMessage)
	router.Handle("message unsubscribe", unsubscribeMessage)

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
