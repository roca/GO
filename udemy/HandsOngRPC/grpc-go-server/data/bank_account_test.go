//go:build unit

// run test with this command: go test . --tags unit --count=1

package data

import "testing"

func TestBankAccount_Table(t *testing.T) {
	s := models.BankAccount.Table()
	if s != "bank_accounts" {
		t.Errorf("Wrong table name returned. Expected 'bank_accounts', got '%s'", s)
	}
}
