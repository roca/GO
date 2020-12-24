package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	// define the book
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var Books = []Book{
	{Title: "The Hitchhikers Guid to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0000000000"},
}

var books = map[string]Book{
	"0345391802": Books[0],
	"0000000000": Books[1],
}

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	bytes, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return bytes
}

// FromJSON to be used for un-marshalling of the Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
// '/api/books'
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	default:
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Unsupported request method."))
		if err != nil {
			log.Println(err)
		}
	}
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
// '/api/books/<isbn>'
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(bytes)
}

func AllBooks() []Book {
	return Books
}

func writeJSON(w http.ResponseWriter, books []Book) {
	bytes, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(bytes)
}
