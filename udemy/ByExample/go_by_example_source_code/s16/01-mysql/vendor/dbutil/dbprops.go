package dbutil

import "fmt"

// Info ...
type Info struct {
	ID   int
	Name string
}

// DbDriver ...
const DbDriver = "mysql"

// User ...
const User = "root"

// Password ...
const Password = "password"

// DbName ...
const DbName = "byexample"

// TableName ...
const TableName = "person"

// DataSourceName ...
// dataSourceName := "root:tyler@tcp(127.0.0.1:3306)/byexample?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(mysqldb:3306)/%s?charset=utf8",
	User, Password, DbName)
