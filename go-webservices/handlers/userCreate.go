package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"encoding/base64"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

//curl http://192.168.99.100:3000/api/users -data "name=nkozyra&first=Nathan&last=Kozyra&email=nathan@nathankozyra.com"

// CreateUser API endpoint
func UserCreate(w http.ResponseWriter, r *http.Request) {

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

	f, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err.Error())
	}

	fileData, _ := ioutil.ReadAll(f)

	fileString := base64.StdEncoding.EncodeToString(fileData)

	// output, err := json.Marshal(NewUser)
	// fmt.Fprintf(w, string(output))

	// if err != nil {
	// 	fmt.Println("Something went wrong")
	// }

	Response := CreateResponse{}

	sql := "INSERT INTO users set " +
		"   user_nickname='" + NewUser.Name +
		"', user_first='" + NewUser.First +
		"', user_last='" + NewUser.Last +
		"', user_email='" + NewUser.Email +
		"', user_image='" + fileString + "'"

	stmtIns, err := db.Prepare(sql) // ? = placeholder
	if err != nil {
		fmt.Println("Could not prepare sql statement")
	}

	q, err := stmtIns.Exec()
	if err != nil {
		errorMessage, errorCode := dbErrorParse(err.Error())
		fmt.Println(errorMessage)
		em := ErrorMessages(errorCode)
		Response.Error = em.Msg
		Response.ErrorCode = em.ErrCode

		//http.Error(w, em.Msg, em.StatusCode)
		fmt.Println(em.StatusCode)
	}

	fmt.Println(q)
	createOutput, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(createOutput))
}
