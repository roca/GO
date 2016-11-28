package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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
