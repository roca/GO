package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/GOCODE/udemy/RESTfulJWT/driver"
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

// postgres://postgres:password@db/jwtexample?options'
func init() {

	db = driver.ConnectDB()

	_, err := db.Exec("DROP TABLE books")
	if err != nil {
		log.Println(err)
	}

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

func signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("successfully called signup"))
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
