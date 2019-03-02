package main

import (
	r "github.com/dancannon/gorethink"
)

func main() {
	r.Connect(r.ConnectOpts{
		Address: "rethinkdb:28015",
	})
}
