package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-oci8"
	"log"
	"time"
)

type Clone struct {
	Id   int `beedb:"PK"`
	Name string
	Seq  string
	//Regions   []int
	CreatedAt time.Time
	CreatedBy string
}

func main() {

	db, err := sql.Open("postgres", "user=keymaster password=keymaster sid=CORELIMSDEV host=renoralimsd port=1531")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	conn_err := db.Ping()
	if conn_err != nil {
		panic(conn_err)
	}

	QueryRows2(db)
}

func QueryRows(db *sql.DB) {
	rows, err := db.Query("SELECT name,seq FROM clones")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name, seq string
		if err := rows.Scan(&name, &seq); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %s\n", name, seq)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
func QueryRows2(db *sql.DB) {

	var orm beedb.Model

	var allclones []Clone

	beedb.PluralizeTableNames = true

	orm = beedb.New(db, "pg")
	beedb.OnDebug = true

	err := orm.FindAll(&allclones)
	if err != nil {
		panic(err)
	}

	for i, clone := range allclones {
		fmt.Printf("row %d : %s %s\n", i, clone.Name, clone.Seq)
	}
}
