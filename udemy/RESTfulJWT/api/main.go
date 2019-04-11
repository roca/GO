package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GOCODE/udemy/RESTfulJWT/api/driver"
	"github.com/gorilla/mux"
)

type User struct {
	IDA      int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {

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
	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleware(protectedEndpoint)).Methods("GET")

	port := 3001

	log.Printf("Listen on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called login"))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called protected"))
}

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleware invoked.")
	return next
}
