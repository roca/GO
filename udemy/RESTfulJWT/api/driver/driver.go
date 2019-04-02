package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	gotenv.Load()

	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRES_CONNECTION_STRING"))
	fmt.Println(pgUrl)
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	log.Println(pgUrl)

	return db
}
