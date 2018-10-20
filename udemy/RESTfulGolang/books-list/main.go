package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []book

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all books is called")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get book is called")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book is called")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update book is called")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove book is called")
}
