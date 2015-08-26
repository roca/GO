package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"

	_ "github.com/mattn/go-oci8"
)

//sqlplus keymaster/keymaster002@//renoralimsd.regeneron.regn.com:1531/CORELIMSDEV
//sqlplus keymaster/keymaster002@//taroralimsp.regeneron.regn.com:1532/CORELIMSPROD

// var orm beedb.Model

type Users struct {
	Id       int
	Ad_Login string
}

type VoxRecord struct {
	VolNo     int64
	Name      string
	MaxVal    int64
	MinVal    int64
	Mean      float64
	StdDev    float64
	Vol_mm3   float64
	RngMean   float64
	RngSDev   float64
	RngVol    float64
	FileName  string
	ObjectMap string
}

type VoxRecords []VoxRecord

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores

	db, err := sql.Open("oci8", "bodycomp_dev/t86fEdbQgKrnL@renoralimsd.regeneron.regn.com:1531/CORELIMSDEV")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	s := bufio.NewScanner(os.Stdin)
	// var wg sync.WaitGroup

	for s.Scan() {
		row := string([]byte(s.Text()))
		columns := regexp.MustCompile("\\s+").Split(row, 13)[1:]
		// wg.Add(1)
		// go func(columns []string) {

		voxr, _ := makeRecord(columns)

		if err := insertRecord(db, voxr); err != nil {
			fmt.Print("----ERROR ENTERING--------", voxr, err, "\n\n")
		}
		fmt.Print("-----ENTERED-------", voxr, "\n\n")
		// 	wg.Done()
		//
		// }(columns)

	}

	// wg.Wait()

	// allUsers, err := selectallUsers(db)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }

	// // orm = beedb.New(db, "oracle")
	// for _, user := range allUsers {
	// 	fmt.Println(user.Id, user.Ad_Login)

	// }
	defer db.Close()

	// if err = testSelect(db); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}

func selectallUsers(db *sql.DB) ([]Users, error) {
	rows, err := db.Query("select id, ad_login from users")
	var allUsers []Users
	if err != nil {
		return allUsers, err
	}
	defer rows.Close()

	for rows.Next() {
		var f1 int
		var f2 string
		rows.Scan(&f1, &f2)
		allUsers = append(allUsers, Users{Id: f1, Ad_Login: f2})
		fmt.Println(f1, f2) // 3.14 foo
	}
	return allUsers, nil
}

func makeRecord(columns []string) (VoxRecord, error) {
	var voxr VoxRecord
	headers := []string{"Vol_#", "Name", "MaxVal", "MinVal", "Mean", "Std.Dev", "Vol_mm3", "RngMean", "RngSDev", "Rng Vol", "File", "Object Map"}
	for i, column := range columns {

		switch i {
		case 0:
			v, err := strconv.ParseInt(column, 0, 64)
			if err == nil {
				voxr.VolNo = v
			}
		case 1:
			voxr.Name = column
		case 2:
			v, err := strconv.ParseInt(column, 0, 64)
			if err == nil {
				voxr.MaxVal = v
			}
		case 3:
			v, err := strconv.ParseInt(column, 0, 64)
			if err == nil {
				voxr.MinVal = v
			}
		case 4:
			v, err := strconv.ParseFloat(column, 64)
			if err == nil {
				voxr.Mean = v
			}
		case 5:
			v, err := strconv.ParseFloat(column, 64)
			if err == nil {
				voxr.StdDev = v
			}
		case 6:
			v, err := strconv.ParseFloat(column, 64)
			if err == nil {
				voxr.Vol_mm3 = v
			}
		case 7:
			v, err := strconv.ParseFloat(column, 64)
			if err == nil {
				voxr.RngMean = v
			}
		case 8:
			v, err := strconv.ParseFloat(column, 64)
			if err == nil {
				voxr.RngSDev = v
			}
		case 9:
			v, err := strconv.ParseFloat(column, 64)
			if err == nil {
				voxr.RngVol = v
			}
		case 10:
			voxr.FileName = column
		case 11:
			voxr.ObjectMap = column
		default:

		}
		fmt.Printf("%d:%20s\t%s\n", i, headers[i], column)
	}
	return voxr, nil

}

func insertRecord(db *sql.DB, voxr VoxRecord) error {

	if voxr.VolNo != 0 {
		_, err := db.Exec("insert into VOXES ( VolNo, Name, MaxVal, MinVal, Mean, StdDev, Vol_mm3, RngMean, RngSDev, RngVol, FileName, ObjectMap ) values(:1,:2,:3,:4,:5,:6,:7,:8,:9,:10,:11,:12)", voxr.VolNo,
			voxr.Name,
			voxr.MaxVal,
			voxr.MinVal,
			voxr.Mean,
			voxr.StdDev,
			voxr.Vol_mm3,
			voxr.RngMean,
			voxr.RngSDev,
			voxr.RngVol,
			voxr.FileName,
			voxr.ObjectMap,
		)
		if err != nil {
			return err
		}
	}

	return nil

}

func testSelect(db *sql.DB) error {
	rows, err := db.Query("select id, ad_login from users")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var f1 float64
		var f2 string
		rows.Scan(&f1, &f2)
		fmt.Println(f1, f2) // 3.14 foo
	}
	_, err = db.Exec("create table foo(bar varchar2(256))")
	_, err = db.Exec("drop table foo")
	if err != nil {
		return err
	}

	return nil
}
