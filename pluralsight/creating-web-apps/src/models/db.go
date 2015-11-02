package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("oci8", "user=pmuser password=pmuser sid=PDB1 host=192.168.59.103 port=1521")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	conn_err := db.Ping()
	if conn_err != nil {
		panic(conn_err)
	}

	return db, err
}
