package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Hello API endpoint
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
