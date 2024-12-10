package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@(127.0.0.1)/testdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database")
}

func main() {
	defer db.Close()

	gRouter := mux.NewRouter()
	gRouter.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":3000", gRouter)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var version string

	if err := db.QueryRow("SELECT VERSION()").Scan(&version); err != nil {
		log.Fatal(err)
	}

	w.Write([]byte("Database version: " + version))
}
