package data

import (
	"fmt"
	"os"	
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	db2 "github.com/upper/db/v4"
)

// Test the New function
func TestNew(t *testing.T) {
	fakeDB, _, _ := sqlmock.New()
	defer fakeDB.Close()

	_ = os.Setenv("DATABASE_TYPE", "postgres")
	m := New(fakeDB)
	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Errorf("Wrong type. Expected data.Models type, got %T", m)
	}

	if fmt.Sprintf("%T", m.BankAccounts) != "data.BankAccount" {
		t.Errorf("Wrong type. Expected data.BankAccount type, got %T", m.BankAccounts)
	}

	if fmt.Sprintf("%T", m.BankTransactions) != "data.BankTransaction" {
		t.Errorf("Wrong type. Expected data.BankTransaction type, got %T", m.BankTransactions)
	}
}

	

// Test the getInsertedID function

func TestGetInsertedID(t *testing.T) {
	var id db2.ID
	id = int64(1)

	returnedID := getInsertedID(id)
	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Errorf("Wrong type. Expected int type, got %T", returnedID)
	}

	id = 1
	returnedID = getInsertedID(id)
	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Errorf("Wrong type. Expected int type, got %T", returnedID)
	}

	id = int32(1)
	returnedID = getInsertedID(id)
	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Errorf("Wrong type. Expected int type, got %T", returnedID)
	}

}
