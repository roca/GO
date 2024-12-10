package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	gRouter := mux.NewRouter()

	http.ListenAndServe(":3000", gRouter)
}
