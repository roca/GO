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

type Model interface {
	Table() string
	GetAll(condition db2.Cond) ([]Model, error)
	Get(id int) (Model, error)
	Update(m Model) error
	Delete(id int) error
	Insert(m Model) (int, error)
	Builder(id int) ([]Model, error)
}

type Models struct {
	// any models inserted here ( and in the New function)
	// are easily accessible from throughout the entire application
	BankAccount
	BankExchangeRate
	BankTransaction
	BankTransfer
}

func New(databasePool *sql.DB) Models {
	db = databasePool

	if !dbCreated {
		switch os.Getenv("DATABASE_TYPE") {
		case "mysql", "mariadb":
			upper, _ = mysql.New(db)
		case "postgres", "postgresql":
			upper, _ = postgresql.New(db)
		default:
			// do nothing
		}
		dbCreated = true
	}

	return Models{
		BankAccount:      BankAccount{},
		BankExchangeRate: BankExchangeRate{},
		BankTransaction:  BankTransaction{},
		BankTransfer:     BankTransfer{},
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
