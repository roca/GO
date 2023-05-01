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

func TestSQLiteRepository_GetHoldingByID(t *testing.T) {
	_ = testRepo.Migrate()
	h := Holdings{
		Amount:        1.0,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	result, _ := testRepo.InsertHoldings(h)
	holding, err := testRepo.GetHoldingByID(result.ID)
	if err != nil {
		t.Error("get holding by id failed", err)
	}

	if holding.ID != result.ID {
		t.Error("invalid holding returned: ", holding.ID)
	}

	_, err = testRepo.GetHoldingByID(0)
	if err == nil {
		t.Error("invalid non-existing holding id should return error")
	}
}

func TestSQLiteRepository_UpdateHolding(t *testing.T) {
	_ = testRepo.Migrate()
	h := Holdings{
		Amount:        1.0,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	result, _ := testRepo.InsertHoldings(h)
	h.Amount = 2.0
	err := testRepo.UpdateHolding(result.ID, h)
	if err != nil {
		t.Error("update holding failed", err)
	}

	holding, _ := testRepo.GetHoldingByID(result.ID)
	if holding.Amount != 2.0 {
		t.Error("invalid holding amount returned: ", holding.Amount)
	}
}

func TestSQLiteRepository_DeleteHolding(t *testing.T) {
	_ = testRepo.Migrate()
	h := Holdings{
		Amount:        1.0,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	result, _ := testRepo.InsertHoldings(h)
	err := testRepo.DeleteHolding(result.ID)
	if err != nil {
		t.Error("delete holding failed", err)
	}

	holding, _ := testRepo.GetHoldingByID(result.ID)
	if holding != nil {
		t.Error("invalid non-existing holding returned: ", holding)
	}
}
