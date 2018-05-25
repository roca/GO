package main

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"fmt"

	"github.com/GOCODE/golang-relay-starter-kit/data"
	"github.com/graphql-go/handler"
)

func main() {

	session, err := mgo.Dial("ec2-52-91-31-26.compute-1.amazonaws.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("quotes")
	count, _ := c.Count()
	fmt.Printf("\n\nConnect to quotes collection you have %d record(s)\n\n", count)

	// simplest relay-compliant graphql server HTTP handler
	h := handler.New(&handler.Config{
		Schema: &data.Schema,
		Pretty: true,
	})

	// create graphql endpoint
	http.Handle("/graphql", h)

	// serve!
	port := ":8080"
	log.Printf(`GraphQL server starting up on http://localhost%v`, port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
