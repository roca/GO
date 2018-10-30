package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/GOCODE/udemy/RESTfulGolang/books-list/models"
	"github.com/GOCODE/udemy/RESTfulGolang/books-list/repository/book"
	"github.com/gorilla/mux"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks ...
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}

		bookRepo := repository.BookRepository{}
		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

// GetBook ...
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		bookRepo := repository.BookRepository{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		book = bookRepo.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

// AddBook ...
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook models.Book
		_ = json.NewDecoder(r.Body).Decode(&newBook)

		bookRepo := repository.BookRepository{}

		newBookID := bookRepo.AddBook(db, newBook)

		log.Println("New book added with ID of", newBookID)

		c.GetBooks(db)(w, r)
	}
}

// UpdateBook ...
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var existingBook models.Book
		_ = json.NewDecoder(r.Body).Decode(&existingBook)

		bookRepo := repository.BookRepository{}

		rowsUpdated := bookRepo.UpdateBook(db, existingBook)

		log.Println("The number of rows affected is", rowsUpdated)

		c.GetBooks(db)(w, r)
	}
}

// RemoveBook ...
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		bookRepo := repository.BookRepository{}

		rowsDeleted := bookRepo.RemoveBook(db, id)

		log.Println("The number of rows deleted is", rowsDeleted)

		c.GetBooks(db)(w, r)
	}
}
