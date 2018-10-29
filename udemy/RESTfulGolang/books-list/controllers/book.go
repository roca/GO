package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/GOCODE/udemy/RESTfulGolang/books-list/models"
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

		rows, err := db.Query("select * from books")
		logFatal(err)

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			logFatal(err)

			books = append(books, book)
		}
		json.NewEncoder(w).Encode(books)
	}
}

// GetBook ...
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		rows := db.QueryRow("select * from books where id=$1", params["id"])

		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		json.NewEncoder(w).Encode(book)
	}
}

// AddBook ...
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook models.Book
		var newBookID int
		_ = json.NewDecoder(r.Body).Decode(&newBook)

		err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
			newBook.Title, newBook.Author, newBook.Year).Scan(&newBookID)
		logFatal(err)

		log.Println("New book added with ID of", newBookID)

		c.GetBooks(db)(w, r)
	}
}

// UpdateBook ...
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook models.Book
		_ = json.NewDecoder(r.Body).Decode(&newBook)

		result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &newBook.Title, &newBook.Author, &newBook.Year, &newBook.ID)

		rowsUpdated, err := result.RowsAffected()
		logFatal(err)

		log.Println("The number of rows affected is", rowsUpdated)

		c.GetBooks(db)(w, r)
	}
}

// RemoveBook ...
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		result, err := db.Exec("delete from books where id=$1", params["id"])
		logFatal(err)

		rowsUpdated, err := result.RowsAffected()
		logFatal(err)

		log.Println("The number of rows deleted is", rowsUpdated)

		c.GetBooks(db)(w, r)
	}
}
