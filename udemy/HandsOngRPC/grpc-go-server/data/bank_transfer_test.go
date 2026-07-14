//go:build unit

// run test with this command: go test . --tags unit --count=1

package data

import "testing"

func TestBankTransfer_Table(t *testing.T) {
	s := models.BankTransfer.Table()
	if s != "bank_transfers" {
		t.Errorf("Wrong table name returned. Expected 'bank_transfers', got '%s'", s)
	}
}
