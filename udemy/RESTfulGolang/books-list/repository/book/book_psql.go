package repository

import (
	"database/sql"
	"log"

	"github.com/GOCODE/udemy/RESTfulGolang/books-list/models"
)

// BookRepository ...
type BookRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks ...
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	return books
}

// GetBook ...
func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {

	rows := db.QueryRow("select * from books where id=$1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	return book
}

// AddBook ..
func (b BookRepository) AddBook(db *sql.DB, book models.Book) int64 {
	return 0
}
