package main

import (
	"net/http"
	"os"

	"github.com/GOCODE/go-webservices/handlers"

	"github.com/gorilla/mux"
)

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", handlers.Hello)
	gorillaRoute.HandleFunc("/api/user/create", handlers.CreateUser).Methods("GET")
	gorillaRoute.HandleFunc("/api/user/read/{id:[0-9]+}", handlers.GetUser).Methods("GET")

	http.Handle("/", gorillaRoute)
	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, nil)
}
