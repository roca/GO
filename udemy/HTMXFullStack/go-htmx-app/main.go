package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var tmpl *template.Template

func init() {

	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}
	log.Println("Connected to the database")

	err = initSchema(db)
	if err != nil {
		log.Fatalf("Failed to initialize schema: %v\n", err)
	}
	log.Println("Initialized the schema")

	tmpl, err = template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v\n", err)
	}
	log.Println("Parsed HTML templates")
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1)/testdb?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v\n", err)
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Database not responding to Pings: %v\n", err)
	}
	return db,nil
}

func initSchema(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id 	INT 		AUTO_INCREMENT PRIMARY KEY NOT NULL, 
		task 	VARCHAR(255)	 NOT NULL, 
		done 	BOOLEAN 	NOT NULL DEFAULT 0
	)
	`)
	if err != nil {
		return fmt.Errorf("Failed to create table: %v", err)
	}

	_,err = db.Exec("TRUNCATE TABLE tasks")
	if err != nil {
		return fmt.Errorf("Failed to truncate task: %v\n", err)
	}

	_,err = db.Exec("INSERT INTO tasks (task) VALUES ('Research the Video')")
	if err != nil {
		return fmt.Errorf("Failed to insert task 1: %v\n", err)
	}

	_,err = db.Exec("INSERT INTO tasks (task) VALUES ('Plan the Video')")
	if err != nil {
		return fmt.Errorf("Failed to insert task 2: %v\n", err)
	}

	_,err = db.Exec("INSERT INTO tasks (task) VALUES ('Record the Video')")
	if err != nil {
		return fmt.Errorf("Failed to insert task 3: %v\n", err)
	}
	return nil
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
