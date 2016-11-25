package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// API struct
type API struct {
	Message string "json:message"
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		message := API{"Hello World"}

		output, err := json.Marshal(message)

		if err != nil {
			fmt.Println("Something went wrong!")
		}

		fmt.Fprintf(w, string(output))

		port := os.Getenv("PORT")

		http.ListenAndServe(":"+port, nil)

	})
}
