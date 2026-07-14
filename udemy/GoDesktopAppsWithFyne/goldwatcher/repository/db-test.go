package repository

import (
	"database/sql"
	"time"
)

type TestRepository struct {
}

func NewTestRepository() Repository {
	return &TestRepository{}
}

func (repo *TestRepository) Connection() *sql.DB {
	return nil
}

func (repo *TestRepository) Migrate() error {
	return nil
}

func (repo *TestRepository) InsertHoldings(h Holdings) (*Holdings, error) {

	return &h, nil
}

func (repo *TestRepository) AllHoldings() ([]Holdings, error) {
	var holdings []Holdings
	h := Holdings{
		Amount:        1.0,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	holdings = append(holdings, h)

	h = Holdings{
		Amount:        2.0,
		PurchaseDate:  time.Now(),
		PurchasePrice: 2000,
	}
	holdings = append(holdings, h)

	return holdings, nil
}

func (repo *TestRepository) GetHoldingByID(id int64) (*Holdings, error) {
	h := Holdings{
		Amount:        1.0,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	return &h, nil
}

func (repo *TestRepository) UpdateHolding(id int64, updated Holdings) error {

	return nil
}

func (repo *TestRepository) DeleteHolding(id int64) error {

	return nil
}
