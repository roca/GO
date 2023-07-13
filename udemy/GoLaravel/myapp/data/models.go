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
var dbCreated bool

type Models struct {
	// any models inserted here ( and in the New function)
	// are easily accessible from throughout the entire application
	Users  User
	Tokens Token
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	if !dbCreated {
		switch os.Getenv("DATABASE_TYPE") {
		case "mysql":
			upper, _ = mysql.New(db)
		case "mariadb":
			upper, _ = mysql.New(db)
		default:
			upper, _ = postgresql.New(db)
		}
		dbCreated = true
	}

	return Models{
		Users:  User{},
		Tokens: Token{},
	}
}

func getInsertedID(i db2.ID) int {
	switch v := i.(type) {
	case int64:
		return int(v)
	case int32:
		return int(v)
	case int:
		return v
	default:
		return v.(int)
	}
}
