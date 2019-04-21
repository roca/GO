package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"udemy.com/RESTfulJWT/api/controllers"
	"udemy.com/RESTfulJWT/api/driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gotenv.Load()

	db = driver.ConnectDB()

	_, err := db.Exec("DROP TABLE users")
	if err != nil {
		log.Println(err)
	}

	createTableSQL := "CREATE TABLE USERS ( "
	createTableSQL += "ID SERIAL PRIMARY KEY,"
	createTableSQL += "EMAIL TEXT NOT NULL UNIQUE,"
	createTableSQL += "PASSWORD TEXT NOT NULL"
	createTableSQL += ")"

	_, err = db.Exec(createTableSQL)
	logFatal(err)

	insertUsersSQL := " INSERT INTO USERS (EMAIL, PASSWORD) VALUES ('test@example.com','12345'); "
	insertUsersSQL += " INSERT INTO USERS (EMAIL, PASSWORD) VALUES ('test123@example.com','abcde');"
	_, err = db.Exec(insertUsersSQL)
	logFatal(err)

}

func main() {
	controller := controllers.Controller{}
	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.Signup(db)).Methods("POST")
	router.HandleFunc("/login", controller.Login(db)).Methods("POST")
	router.HandleFunc("/protected", controller.TokenVerifyMiddleware(controller.ProtectedEndpoint())).Methods("GET")

	port := 3001

	log.Printf("Listen on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
