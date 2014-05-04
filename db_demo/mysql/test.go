package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
	"time"
)

type Command struct {
	Id   int `beedb:"PK"`
	Path string
	Dir  string
	//Regions   []int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func main() {

	db, err := sql.Open("mymysql", "commandlogs_dev/commander/cody")
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
	rows, err := db.Query("SELECT path,dir FROM commands")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var path, dir string
		if err := rows.Scan(&path, &dir); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %s\n", path, dir)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
func QueryRows2(db *sql.DB) {

	var orm beedb.Model

	var allCommands []Command

	beedb.PluralizeTableNames = true

	orm = beedb.New(db, "pg")
	beedb.OnDebug = true

	err := orm.FindAll(&allCommands)
	if err != nil {
		panic(err)
	}

	for i, command := range allCommands {
		fmt.Printf("row %d : %s %s\n", i, command.Path, command.Dir)
	}
}
