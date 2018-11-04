package main

import (
	"database/sql"
	"fmt"

	dbu "./dbutil"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {

	fmt.Printf("Opening the %s ... ", dbu.DbName)
	db, err = sql.Open(dbu.DbDriver, dbu.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success!")
	}

	defer db.Close()

	// test1Select()
	// test2Select()
	// test3Insert()
	// test4Update()
	// test5Delete()
}

func test1Select() {
	fmt.Println("\n==> test1Select")

	// Execute the query
	results, err := db.Query("SELECT id, name FROM " + dbu.TableName)
	checkErr(err)

	for results.Next() {
		var info dbu.Info

		err = results.Scan(&info.ID, &info.Name)
		checkErr(err)

		fmt.Printf("%d\t%s \n", info.ID, info.Name)
	}
}

func test2Select() {
	fmt.Println("\n==> test2Select")
	var info dbu.Info

	var sqlStmt = fmt.Sprintf("SELECT id, name FROM %s where id = ?", dbu.TableName)

	err = db.QueryRow(sqlStmt, 3).Scan(&info.ID, &info.Name)
	checkErr(err)

	fmt.Printf("%d\t%s \n", info.ID, info.Name)
}

func test3Insert() {
	fmt.Println("\n==> test3Insert")

	var sqlStmt = fmt.Sprintf("INSERT %s SET id=?,name=?", dbu.TableName)
	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	res, err := stmt.Exec(5, "George")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func test4Update() {
	fmt.Println("\n==> test4Update")

	var sqlStmt = fmt.Sprintf("update %s set name=? where id=?", dbu.TableName)
	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	res, err := stmt.Exec("Eric", 5)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func test5Delete() {
	fmt.Println("\n==> test5Delete")

	var sqlStmt = fmt.Sprintf("delete from %s where id=?", dbu.TableName)
	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	res, err := stmt.Exec(5)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
