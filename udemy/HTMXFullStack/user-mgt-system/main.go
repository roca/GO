package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var tmpl *template.Template

func init() {
	var err error

	db, err = initDB()
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
	db, err := sql.Open("mysql", "root:root@(127.0.0.1)/usermanagement?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v\n", err)
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Database not responding to Pings: %v\n", err)
	}
	return db, nil
}

func initSchema(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
		id 		INT 		AUTO_INCREMENT PRIMARY KEY NOT NULL, 
		email 		VARCHAR(255)	 NOT NULL, 
		password 	VARCHAR(255)	 NOT NULL, 
		name 		VARCHAR(255)	 NULL, 
		category 	INT		 NULL,
		dob 		DATE		 NULL,
		bio 		LONGTEXT	 NULL,
		avatar 		VARCHAR(255)	 NULL 
	)
	`)
	if err != nil {
		return fmt.Errorf("Failed to create table: %v", err)
	}

	_, err = db.Exec("TRUNCATE TABLE users")
	if err != nil {
		return fmt.Errorf("Failed to truncate task: %v\n", err)
	}

	return nil
}

func main() {
	defer db.Close()
}
