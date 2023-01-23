package main

import (
	"log"
	"net/http"

	"github.com/roca/go-toolkit"
)

type RequestPayload struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ResponsePayload struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code,omitempty"`
}

func main() {
	mux := routes()

	log.Println("Starting server on port 8081")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {

	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/receive-post", receivePost)
	mux.HandleFunc("/remote-service", remoteService)
	mux.HandleFunc("/simulated-service", simulatedService)

	return mux
}

func receivePost(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message:    "hit the handler okay, and sending back a response",
		StatusCode: http.StatusOK,
	}

	err = t.WriteJSON(w, http.StatusAccepted, responsePayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}
}

func remoteService(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	_, statusCode, err := t.PushJSONToRemote("http://localhost:8081/simulated-service", requestPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message:    "hit the handler okay, and sending back a response",
		StatusCode: statusCode,
	}

	err = t.WriteJSON(w, http.StatusAccepted, responsePayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}
}

func simulatedService(w http.ResponseWriter, r *http.Request) {
	payload := ResponsePayload{
		Message:    "simulated service",
	}

	var t toolkit.Tools
	_ = t.WriteJSON(w, http.StatusOK, payload)
}
