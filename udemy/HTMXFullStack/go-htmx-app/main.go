package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var tmpl *template.Template

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@(127.0.0.1)/testdb?parseTime=true")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Database not responding to Pings: %v\n", err)
	}
	log.Println("Connected to the database")

	tmpl, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v\n", err)
	}
	log.Println("Parsed HTML templates")
}

func main() {
	defer db.Close()

	gRouter := mux.NewRouter()
	gRouter.HandleFunc("/", HomeHandler)

	log.Println("Server started on http://localhost:3000")
	http.ListenAndServe(":3000", gRouter)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
