package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/GOCODE/udemy/RESTfulGolang/books-list/controllers"
	"github.com/GOCODE/udemy/RESTfulGolang/books-list/driver"

	"github.com/GOCODE/udemy/RESTfulGolang/books-list/models"

	"github.com/gorilla/mux"
)

var books []models.Book
var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {

	db = driver.ConnectDB()

	_, err := db.Exec("DROP TABLE books")
	if err != nil {
		log.Println(err)
	}

	createTableSQL := "CREATE TABLE books ( "
	createTableSQL += "ID     SERIAL,"
	createTableSQL += "Title  character varying NOT NULL,"
	createTableSQL += "Author character varying NOT NULL,"
	createTableSQL += "Year   character varying NOT NULL"
	createTableSQL += ")"

	_, err = db.Exec(createTableSQL)
	logFatal(err)

	insertBooksSQL := " INSERT INTO books (title, author, year) VALUES ('Golang pointers', 'Mr. Golang', '2010' ); "
	insertBooksSQL += " INSERT INTO books (title, author, year) VALUES ('Goroutines', 'Mr. Goroutine', '2011' ); "
	insertBooksSQL += " INSERT INTO books (title, author, year) VALUES ('Golang routers', 'Mr. Router', '2012' ); "
	insertBooksSQL += " INSERT INTO books (title, author, year) VALUES ('Golang concurrency', 'Mr. Currency', '2013' ); "
	insertBooksSQL += " INSERT INTO books (title, author, year) VALUES ('Golang good parts', 'Mr. Good', '2014' ); "
	_, err = db.Exec(insertBooksSQL)
	logFatal(err)
}
func main() {

	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	// router.HandleFunc("/books", updateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/")))
	router.PathPrefix("/swagger-ui/").Handler(sh)

	log.Fatal(http.ListenAndServe(":8000", router))

}
