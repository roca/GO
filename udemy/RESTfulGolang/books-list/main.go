package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

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

	rows := db.QueryRow("select * from books where id=$1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook Book
	var newBookID int
	_ = json.NewDecoder(r.Body).Decode(&newBook)

	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		newBook.Title, newBook.Author, newBook.Year).Scan(&newBookID)
	logFatal(err)

	log.Println("New book added with ID of", newBookID)

	getBooks(w, r)

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

	_, err := db.Exec("delete from books where id=$1", params["id"])
	logFatal(err)

	getBooks(w, r)
}
