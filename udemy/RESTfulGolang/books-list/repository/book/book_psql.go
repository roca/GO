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
	var newBookID int64

	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&newBookID)
	logFatal(err)

	return newBookID
}

// UpdateBook ...
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {

	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	return rowsUpdated

}

// RemoveBook ...
func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {

	result, err := db.Exec("delete from books where id=$1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	log.Println("The number of rows deleted is", rowsDeleted)

	return rowsDeleted
}
