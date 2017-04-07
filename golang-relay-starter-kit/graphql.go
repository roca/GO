package main

import (
	"log"
	"net/http"

	"github.com/GOCODE/golang-relay-starter-kit/data"
	"github.com/graphql-go/handler"
)

func main() {

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
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ListenAndServe failed, %v", err)
	}
}
