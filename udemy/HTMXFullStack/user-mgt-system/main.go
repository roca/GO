package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"user-mgt-system/pkg/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var db *sql.DB
var tmpl *template.Template
var Store = sessions.NewCookieStore([]byte("usermanagementsecret"))

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

	// Set up Sessions
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 3,
		HttpOnly: true,
	}
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
		id 		CHAR(36) 	 PRIMARY KEY NOT NULL, 
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

	// _, err = db.Exec("TRUNCATE TABLE users")
	// if err != nil {
	// 	return fmt.Errorf("Failed to truncate task: %v\n", err)
	// }

	return nil
}

func main() {
	defer db.Close()

	gRouter := mux.NewRouter()

	// Setup Static file handling for images

	fileServer := http.FileServer(http.Dir("./uploads"))
	gRouter.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads", fileServer))

	//All dynamic routes
	
	gRouter.HandleFunc("/", handlers.Homepage(db, tmpl, Store)).Methods("GET")

	gRouter.HandleFunc("/register", handlers.RegisterPage(db, tmpl)).Methods("GET")
	gRouter.HandleFunc("/register", handlers.RegisterHandler(db, tmpl)).Methods("POST")

	gRouter.HandleFunc("/login", handlers.LoginPage(db, tmpl)).Methods("GET")
	gRouter.HandleFunc("/login", handlers.LoginHandler(db, tmpl, Store)).Methods("POST")

	gRouter.HandleFunc("/edit", handlers.Editpage(db,tmpl,Store)).Methods("GET")
	gRouter.HandleFunc("/edit", handlers.UpdateProfileHandler(db,tmpl,Store)).Methods("POST")

	gRouter.HandleFunc("/upload-avatar", handlers.AvatarPage(db,tmpl,Store)).Methods("GET")
	gRouter.HandleFunc("/upload-avatar", handlers.UploadAvatarHandler(db,tmpl,Store)).Methods("POST")

	gRouter.HandleFunc("/logout", handlers.LogoutHandler(Store)).Methods("GET")

	log.Println("Server started on http://localhost:4000")
	http.ListenAndServe(":4000", gRouter)
}
