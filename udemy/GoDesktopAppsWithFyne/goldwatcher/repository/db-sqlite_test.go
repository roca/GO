package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed", err)
	}
}

func TestSQLiteRepository_InsertHoldings(t *testing.T) {
	_ = testRepo.Migrate()
	h := Holdings{
		Amount:        1.0,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	result, err := testRepo.InsertHoldings(h)
	if err != nil {
		t.Error("insert failed", err)
	}

	if result.ID <= 0 {
		t.Error("invalid id sent back: ", err)
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	_ = testRepo.Migrate()
	h := Holdings{
		Amount:        1.0,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	_, _ = testRepo.InsertHoldings(h)
	holdings, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("all holdings failed", err)
	}

	if len(holdings) < 1 {
		t.Error("invalid number of holdings returned: ", len(holdings))
	}
}
