package dbutil

import "fmt"

// Info ...
type Info struct {
	ID   int
	Name string
}

// DBMode ...
// **added for PostgreSQL** - choose "postgres" or "mysql"!
const DBMode = "mysql"

// const DBMode = "postgres"

// DbDriver ...
// const DbDriver = "mysql"
var DbDriver = "mysql" // **added for PostgreSQL**

// User ...
// const User = "root"
var User = "postgres" // **added for PostgreSQL**

// Password ...
const Password = "postgres"

// DbName ...
const DbName = "byexample"

// TableName ...
const TableName = "person"

// DataSourceName ...
// dataSourceName := "root:tyler@tcp(postgresdb:3306)/byexample?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(postgresdb:3306)/%s?charset=utf8",
	User, Password, DbName)
