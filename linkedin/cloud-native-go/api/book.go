package api

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	// define the book
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
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

var Books = []Book{
	{Title: "The Hitchhikers Guid to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0000000000"},
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(Books)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	w.Write(bytes)
}
