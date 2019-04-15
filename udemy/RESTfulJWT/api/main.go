package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"udemy.com/RESTfulJWT/api/driver"

	"github.com/gorilla/mux"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       int    `json:"id"`
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

func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = string(hash)

	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"

	err = db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		error.Message = "Server error."
		respondWithError(w, http.StatusInternalServerError, error)
		return
	}

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	responseJSON(w, user)

}

func GenerateToken(user User) (string, error) {
	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func login(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	token, err := GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)
	//w.Write([]byte("successfully called login"))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called protected"))
}

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fmt.Println("TokenVerifyMiddleware invoked.")
	return next
}
