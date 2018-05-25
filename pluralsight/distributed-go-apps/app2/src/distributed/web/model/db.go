package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("postgres", "postgres://developer:dbpass@192.168.33.12/distributed")
	if err != nil {
		panic(err.Error())
	}
}
