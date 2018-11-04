package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	utl "./util"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	maxCPU := runtime.NumCPU()

	utl.CPUUsed = 4
	runtime.GOMAXPROCS(utl.CPUUsed)

	fmt.Printf("\n= = = = = = = = = = = = = = = = = = = = = = = = = = =\n")
	fmt.Printf("= Number of CPUs (Total=%d - Used=%d) \n", maxCPU, utl.CPUUsed)
}

func main() {

	fmt.Printf("= Opening the %s ... ", utl.DbName)
	var err error
	utl.DB, err = sql.Open(utl.DbDriver, utl.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Success!")
		fmt.Printf("= = = = = = = = = = = = = = = = = = = = = = = = = = =\n\n")
	}

	defer utl.DB.Close()

	utl.MatchCustomersHouses()
	utl.LogMrtgApps()
}
