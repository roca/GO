package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleware(protectedEndpoint)).Methods("GET")

	port := 3001

	log.Printf("Listen on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup invoked.")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login invoked.")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint invoked.")
}

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleware invoked.")
	return nil
}
