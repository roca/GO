package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// API struct
type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	name := urlParams["user"]

	HelloMessage := "Hello, " + name
	message := API{HelloMessage}

	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output))

}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	http.Handle("/", gorillaRoute)
	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, nil)
}
