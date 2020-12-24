package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
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
	if err := json.Unmarshal(data, &book); err != nil {
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
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Unsupported request method.")); err != nil {
			log.Println(err)
		}
	}
}

// BooksHandleFunc to be used as http.HandleFunc for Book API
// '/api/books/<isbn>'
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := path.Base(r.URL.Path)
	//fmt.Println(isbn)
	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("Unsupported request method.")); err != nil {
			log.Println(err)
		}
	}
}

func AllBooks() []Book {
	return Books
}

func CreateBook(book Book) (string, bool) {
	if b, exists := books[book.ISBN]; exists {
		return b.ISBN, false
	}
	Books = append(Books, book)
	books[book.ISBN] = book
	return book.ISBN, true
}

func GetBook(isbn string) (Book, bool) {
	if b, exists := books[isbn]; exists {
		//fmt.Println(b)
		return b, true
	}
	return Book{}, false
}

func UpdateBook(isbn string, book Book) bool {
	for i := 0; i < len(Books); i++ {
		if _, exists := books[Books[i].ISBN]; exists && Books[i].ISBN == isbn {
			Books[i].Title = book.Title
			Books[i].Author = book.Author
			Books[i].Description = book.Description
			return true
		}
	}

	return false
}

func DeleteBook(isbn string) {
	for i := 0; i < len(Books); i++ {
		if _, exists := books[isbn]; exists && Books[i].ISBN == isbn {
			delete(books,isbn)
			Books = append(Books[0:i], Books[i+1:]...)
		}
	}	
}

func writeJSON(w http.ResponseWriter, value interface{}) {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset-utf-8")
	if _, err := w.Write(bytes); err != nil {
		log.Println(err)
	}
}
