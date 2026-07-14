package main

import (
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
)

/*

Channel Events
    'channel add'
    'channel subscribe'
    'channel unsubscribe'

User events
    'user edit'
    'user subscribe'
    'user unsubscribe'

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

	//Message Events
	router.Handle("message add", addChannelMessage)
	router.Handle("message subscribe", subscribeChannelMessage)
	router.Handle("message unsubscribe", unsubscribeChannelMessage)

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
