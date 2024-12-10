package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	gRouter := mux.NewRouter()
	gRouter.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", gRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page"))
}
