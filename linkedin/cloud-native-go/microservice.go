package main

import (
	"linkedin/cloud-native-go/api"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", api.Index)
	http.HandleFunc("/api/echo", api.Echo)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

