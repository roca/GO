package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-oci8"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("oci8", "pmuser/pmuser@192.168.59.103:1521/PDB1")
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
