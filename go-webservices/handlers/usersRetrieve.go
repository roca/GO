package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//curl http://192.168.99.100:3000/api/users

// GetUser API endpoint
func UsersRetrieve(w http.ResponseWriter, r *http.Request) {

	log.Println("starting retrieval")
	start := 0
	limit := 10

	next := start + limit
	link := fmt.Sprintf("<http://localhost:8080/api/users?start=%d; rel=\"next\"", next)

	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Link", link)

	password := os.Getenv("MYSQL_ROOT_PASSWORD")

	db, err := sql.Open("mysql", "root:"+password+"@tcp(mysql:3306)/social_network")
	if err != nil {
		fmt.Println("Could not connect to the database")
	}

	rows, _ := db.Query("select * from users LIMIT 10")
	Response := Users{}

	for rows.Next() {
		user := User{}
		rows.Scan(
			&user.ID,
			&user.Name,
			&user.First,
			&user.Last,
			&user.Email,
		)
		Response.Users = append(Response.Users, user)
	}

	output, _ := json.Marshal(Response)
	fmt.Fprintf(w, string(output))
}
