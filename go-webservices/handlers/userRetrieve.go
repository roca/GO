package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

//curl http://192.168.99.100:3000/api/user/read/111

// GetUser API endpoint
func UserRetrieve(w http.ResponseWriter, r *http.Request) {

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
