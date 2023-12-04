//go:build unit
package data

import (
	"fmt"
	"testing"

	db2 "github.com/upper/db/v4"
)

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
