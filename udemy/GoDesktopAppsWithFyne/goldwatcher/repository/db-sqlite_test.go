package repository

import "testing"

func TestSQLiteRepository(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed", err)
	}
}
