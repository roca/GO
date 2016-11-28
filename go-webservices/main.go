package main

import (
	"net/http"
	"os"

	"github.com/GOCODE/go-webservices/handlers"
	"github.com/gorilla/mux"
)

//http://192.168.99.100:3000/api/user/create?user=nkozyra&first=Nathan&last=Kozyra&email=nathan@nathankozyra.com

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", handlers.Hello)
	gorillaRoute.HandleFunc("/api/user/create", handlers.CreateUser).Methods("GET")

	http.Handle("/", gorillaRoute)
	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, nil)
}
