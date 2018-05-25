package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-oci8"
)

type User struct {
	ID       int `beedb:"PK"`
	AD_LOGIN string
	EMAIL    string
}

func main() {

	db, err := sql.Open("oci8", "keymaster/keymaster002@taroralimsp.regeneron.regn.com:1532/CORELIMSPROD")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	conn_err := db.Ping()
	if conn_err != nil {
		panic(conn_err)
	}

	QueryRows(db)
}

func QueryRows(db *sql.DB) {
	rows, err := db.Query("SELECT AD_LOGIN,EMAIL FROM USERS")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var AD_LOGIN, EMAIL string
		if err := rows.Scan(&AD_LOGIN, &EMAIL); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %s\n", AD_LOGIN, EMAIL)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

// func QueryRows2(db *sql.DB) {
//
// 	var orm beedb.Model
//
// 	var allclones []Clone
//
// 	beedb.PluralizeTableNames = true
//
// 	orm = beedb.New(db, "pg")
// 	beedb.OnDebug = true
//
// 	err := orm.FindAll(&allclones)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	for i, clone := range allclones {
// 		fmt.Printf("row %d : %s %s\n", i, clone.Name, clone.Seq)
// 	}
// }
