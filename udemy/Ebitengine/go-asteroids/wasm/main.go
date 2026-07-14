package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", http.FileServer(http.Dir("./")))
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
