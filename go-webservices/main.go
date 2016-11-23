package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received Request: ", r.URL.Path)
	fmt.Fprintf(w, "Hello World!")
}
func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Could not listen: ", err)
	}
}
