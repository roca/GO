package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"database/sql"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book
var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gotenv.Load()

	//Make sure you setup the ELEPHANTSQL_URL to be a uri, e.g. 'postgres://user:pass@host/db?options'
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	log.Println(pgUrl)

	_, err = db.Exec("DROP TABLE books")
	logFatal(err)

	createTableSQL := "CREATE TABLE books ( "
	createTableSQL += "ID     integer NOT NULL,"
	createTableSQL += "Title  character varying NOT NULL,"
	createTableSQL += "Author character varying NOT NULL,"
	createTableSQL += "Year   character varying NOT NULL"
	createTableSQL += ")"

	_, err = db.Exec(createTableSQL)
	logFatal(err)

	insertBooksSQL := " INSERT INTO books (id, title, author, year) VALUES ( 1, 'Golang pointers', 'Mr. Golang', '2010' ); "
	insertBooksSQL += " INSERT INTO books (id, title, author, year) VALUES ( 2, 'Goroutines', 'Mr. Goroutine', '2011' ); "
	insertBooksSQL += " INSERT INTO books (id, title, author, year) VALUES ( 3, 'Golang routers', 'Mr. Router', '2012' ); "
	insertBooksSQL += " INSERT INTO books (id, title, author, year) VALUES ( 4, 'Golang concurrency', 'Mr. Currency', '2013' ); "
	insertBooksSQL += " INSERT INTO books (id, title, author, year) VALUES ( 5, 'Golang good parts', 'Mr. Good', '2014' ); "
	_, err = db.Exec(insertBooksSQL)
	logFatal(err)
}
func main() {

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/")))
	router.PathPrefix("/swagger-ui/").Handler(sh)

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book Book
	books = []Book{}

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

func getBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	params := mux.Vars(r)

	rows, err := db.Query("select * from books where id=$1", params["id"])
	logFatal(err)

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
	}

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)

	books = append(books, newBook)
	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	_ = json.NewDecoder(r.Body).Decode(&newBook)

	for i, book := range books {
		if book.ID == newBook.ID {
			books[i] = newBook
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}