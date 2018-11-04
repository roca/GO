package main

import (
	"database/sql"
	"fmt"

	dbu "./dbutil"
	_ "github.com/go-sql-driver/mysql" // $ go get github.com/go-sql-driver/mysql
	// **added for PostgreSQL**
	_ "github.com/lib/pq" // $ go get github.com/lib/pq
)

var db *sql.DB
var err error

func init() { // **added for PostgreSQL**

	if dbu.DBMode == "postgres" {
		dbu.DbDriver = "postgres"
		dbu.User = "postgres"

		dbu.DataSourceName = fmt.Sprintf("host=localhost port=5432 user=%s "+
			"password=%s dbname=%s sslmode=disable", dbu.User, dbu.Password, dbu.DbName)
	}
}

func main() {

	fmt.Printf("Opening the %s ... ", dbu.DbName)
	db, err = sql.Open(dbu.DbDriver, dbu.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success!")
	}

	defer db.Close()

	test1Select()
	test2Select()
	test3Insert()
	test4Update()
	test5Delete()
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

	// **added for PostgreSQL**
	var sqlStmt string
	if dbu.DBMode == "mysql" {
		sqlStmt = fmt.Sprintf("SELECT id, name FROM %s where id = ?", dbu.TableName)
	} else {
		sqlStmt = fmt.Sprintf("SELECT id, name FROM %s where id = $1", dbu.TableName)
	}

	err = db.QueryRow(sqlStmt, 3).Scan(&info.ID, &info.Name)
	checkErr(err)

	fmt.Printf("%d\t%s \n", info.ID, info.Name)
}

func test3Insert() {
	fmt.Println("\n==> test3Insert")

	// **added for PostgreSQL**
	var sqlStmt string
	if dbu.DBMode == "mysql" {
		sqlStmt = fmt.Sprintf("INSERT %s SET id=?,name=?", dbu.TableName)
	} else {
		sqlStmt = fmt.Sprintf("INSERT INTO %s VALUES ($1,$2)", dbu.TableName)
	}

	stmt, err := db.Prepare(sqlStmt)
	checkErr(err)

	var res sql.Result // **added for PostgreSQL**
	res, err = stmt.Exec(5, "George")
	checkErr(err)

	// **added for PostgreSQL** - LastInsertId() doesn't work with 'PostgreSQL'
	if dbu.DBMode != "postgres" {
		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
	}
}

func test4Update() {
	fmt.Println("\n==> test4Update")

	// **added for PostgreSQL**
	var sqlStmt string
	if dbu.DBMode == "mysql" {
		sqlStmt = fmt.Sprintf("update %s set name=? where id=?", dbu.TableName)
	} else {
		sqlStmt = fmt.Sprintf("update %s set name=$1 where id=$2", dbu.TableName)
	}

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

	// **added for PostgreSQL**
	var sqlStmt string
	if dbu.DBMode == "mysql" {
		sqlStmt = fmt.Sprintf("delete from %s where id=?", dbu.TableName)
	} else {
		sqlStmt = fmt.Sprintf("delete from %s where id=$1", dbu.TableName)
	}

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
