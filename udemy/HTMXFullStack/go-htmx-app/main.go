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

type Task struct {
	ID   int
	Task string
	Done bool
}

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
	db, err := sql.Open("mysql", "root:root@(127.0.0.1)/testdb?parseTime=true")
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
	CREATE TABLE IF NOT EXISTS tasks (
		id 	INT 		AUTO_INCREMENT PRIMARY KEY NOT NULL, 
		task 	VARCHAR(255)	 NOT NULL, 
		done 	BOOLEAN 	NOT NULL DEFAULT 0
	)
	`)
	if err != nil {
		return fmt.Errorf("Failed to create table: %v", err)
	}

	_, err = db.Exec("TRUNCATE TABLE tasks")
	if err != nil {
		return fmt.Errorf("Failed to truncate task: %v\n", err)
	}

	_, err = db.Exec("INSERT INTO tasks (task) VALUES ('Research the Video')")
	if err != nil {
		return fmt.Errorf("Failed to insert task 1: %v\n", err)
	}

	_, err = db.Exec("INSERT INTO tasks (task) VALUES ('Plan the Video')")
	if err != nil {
		return fmt.Errorf("Failed to insert task 2: %v\n", err)
	}

	_, err = db.Exec("INSERT INTO tasks (task) VALUES ('Record the Video')")
	if err != nil {
		return fmt.Errorf("Failed to insert task 3: %v\n", err)
	}
	return nil
}

func main() {
	defer db.Close()

	gRouter := mux.NewRouter()
	gRouter.HandleFunc("/", HomeHandler)

	//Get Tasks
	gRouter.HandleFunc("/tasks", fetchTasks).Methods("GET")

	log.Println("Server started on http://localhost:3000")
	http.ListenAndServe(":3000", gRouter)
}

func fetchTasks(w http.ResponseWriter, r *http.Request) {
	tasks, _ := getTasks(db)
	err := tmpl.ExecuteTemplate(w, "todoList", tasks)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func getTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query("SELECT id, task, done FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("Failed to query tasks: %v\n", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Task, &t.Done)
		if err != nil {
			return nil, fmt.Errorf("Failed to scan task: %v\n", err)
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func getTaskByID(db *sql.DB, id int) (Task, error) {
	row := db.QueryRow("SELECT id, task, done FROM tasks WHERE id = ?", id)
	var t Task
	err := row.Scan(&t.ID, &t.Task, &t.Done)
	if err != nil {
		return Task{}, fmt.Errorf("Failed to scan task: %v\n", err)
	}
	return t, nil
}
