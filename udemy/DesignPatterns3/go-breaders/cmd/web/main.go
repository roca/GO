package main

import (
	"fmt"
	"net/http"
)

const port = ":4000"

type application struct{}

func main() {
	app := application{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	http.ListenAndServe(port, nil)

}
