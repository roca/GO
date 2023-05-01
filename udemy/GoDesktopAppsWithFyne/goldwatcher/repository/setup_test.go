package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/glebarez/go-sqlite"
)

var testRepo *SQLiteRepository
var path = "./testdata/sql.db"

func setup() {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatal(err)
	}

	testRepo = NewSQLiteRepository(db)
}

func teardown() {
	_ = os.Remove(path)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
