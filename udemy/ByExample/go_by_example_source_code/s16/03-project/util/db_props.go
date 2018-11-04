package util

import (
	"database/sql"
	"fmt"
)

// Customer ...
type Customer struct {
	CID         uint    `json:"id"` //Customer ID
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	CreditScore uint    `json:"credit_score"` // 0-800 (must be >650)
	Salary      float32 `json:"salary"`
	Downpayment float32 `json:"downpayment"` //Actual Downpayment
	HouseID     uint    `json:"house_id"`    //Conncects 'Customer' and 'House'
}

// House ...
type House struct {
	HID             uint    `json:"id"`    // House ID
	Price           float32 `json:"price"` // House Price
	MinDownpayment  float32 `json:"min_down"`
	PropertyTax     float32 `json:"prop_tax"` // Monthly
	MaintenanceCost float32 `json:"m_cost"`   // Monthly
}

// Mortgage ...
type Mortgage struct {
	Customer
	House
}

// DbDriver ...
const DbDriver = "mysql"

// User ...
const User = "root"

// Password ...
const Password = "tyler"

// DbName ...
const DbName = "go_db1"

// TableName1 ...
const TableName1 = "customer"

// TableName2 ...
const TableName2 = "house"

// DataSourceName ...
// dataSourceName := "root:tyler@tcp(127.0.0.1:3306)/gwp?charset=utf8"
var DataSourceName = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8",
	User, Password, DbName)

// DB ...
var DB *sql.DB

// RowsCustomer ...
var RowsCustomer *sql.Rows

// NumOfRowsCustomer ...
var NumOfRowsCustomer int

// Mrtgs ...
var Mrtgs []Mortgage

// AppprovedIdx ... indices of Mrtgs[] those indicate approved mortgage applications
var AppprovedIdx []int

// RejectedIdx ... indices of Mrtgs[] those indicate rejected mortgage applications
var RejectedIdx []int
