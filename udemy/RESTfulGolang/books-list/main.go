package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "database/sql"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gotenv.Load()
}
func main() {
	//Make sure you setup the ELEPHANTSQL_URL to be a uri, e.g. 'postgres://user:pass@host/db?options'
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	log.Println(pgUrl)

	// router := mux.NewRouter()

	// books = append(books,
	// 	book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
	// 	book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
	// 	book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
	// 	book{ID: 4, Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
	// 	book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"},
	// )

	// router.HandleFunc("/books", getBooks).Methods("GET")
	// router.HandleFunc("/books/{id}", getBook).Methods("GET")
	// router.HandleFunc("/books", addBook).Methods("POST")
	// router.HandleFunc("/books", updateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

<<<<<<< HEAD
	// log.Fatal(http.ListenAndServe(":8000", router))
=======
	sh := http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/")))
	router.PathPrefix("/swagger-ui/").Handler(sh)

	log.Fatal(http.ListenAndServe(":8000", router))
>>>>>>> 51221302b96aae054eb2f8f222bb5352aeae7fe5

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(&book)
			break
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newBook book
	_ = json.NewDecoder(r.Body).Decode(&newBook)

	books = append(books, newBook)

	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var newBook book
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
