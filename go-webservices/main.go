package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// API struct
type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	name := urlParams["user"]

	HelloMessage := "Hello, " + name
	message := API{HelloMessage}

	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output))

}

type User struct {
	ID    int    "json:id"
	Name  string "json:username"
	Email string "json:email"
	First string "json:first"
	Last  string "json:last"
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	password := os.Getenv("MYSQL_ROOT_PASSWORD")

	db, err := sql.Open("mysql", "root:"+password+"@tcp(mysql:3306)/social_network")
	if err != nil {
		fmt.Println("Could not connect to the database")
	}

	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")

	output, err := json.Marshal(NewUser)
	fmt.Fprintf(w, string(output))

	if err != nil {
		fmt.Println("Something went wrong")
	}

	sql := "INSERT INTO users set " +
		"   user_nickname='" + NewUser.Name +
		"', user_first='" + NewUser.First +
		"', user_last='" + NewUser.Last +
		"', user_email='" + NewUser.Email + "'"

	stmtIns, err := db.Prepare(sql) // ? = placeholder
	if err != nil {
		fmt.Println("Could not prepare sql statement")
	}

	q, err := stmtIns.Exec()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

//http://192.168.99.100:3000/api/user/create?user=nkozyra&first=Nathan&last=Kozyra&email=nathan@nathankozyra.com
func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Pragma", "no-cache")

	urlParams := mux.Vars(r)
	id := urlParams["id"]
	ReadUser := User{}

	password := os.Getenv("MYSQL_ROOT_PASSWORD")

	db, err := sql.Open("mysql", "root:"+password+"@tcp(mysql:3306)/social_network")
	if err != nil {
		fmt.Println("Could not connect to the database")
	}

	er := db.QueryRow("select * from users where user_id = ?", id).Scan(
		&ReadUser.ID,
		&ReadUser.Name,
		&ReadUser.First,
		&ReadUser.Last,
		&ReadUser.Email,
	)

	switch {
	case er == sql.ErrNoRows:
		fmt.Fprintf(w, "No such user")
	case er != nil:
		log.Fatal(er)
		fmt.Fprintf(w, "ERROR")
	default:
		output, _ := json.Marshal(ReadUser)
		fmt.Fprintf(w, string(output))
	}
}

//curl http://192.168.99.100:3000/api/user/read/111

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	gorillaRoute.HandleFunc("/api/user/create", CreateUser).Methods("GET")
	gorillaRoute.HandleFunc("/api/user/read/{id:[0-9]+}", GetUser).Methods("GET")

	http.Handle("/", gorillaRoute)
	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, nil)
}
