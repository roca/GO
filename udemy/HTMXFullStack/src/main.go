package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to some home route")
	})

	http.ListenAndServe(":3000", nil)
}
