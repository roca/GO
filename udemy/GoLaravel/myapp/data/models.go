package data

import (
	"database/sql"
	"os"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper db2.Session

type Models struct {
	// any models inserted here ( and in the New function)
	// are easily accessible from throughout the entire application
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	if os.Getenv("DATABASE_TYPE") == "mysql " || os.Getenv("DATABASE_TYPE") == "mariadb" {
		// TODO: add mysql models
		upper, _ = mysql.New(db)
	} else {
		upper, _ = postgresql.New(db)
	}

	return Models{}
}
