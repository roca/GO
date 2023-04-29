package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{Conn: db}
}

func (repo *SQLiteRepository) Connection() *sql.DB {
	return repo.Conn
}

func (repo *SQLiteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS holdings (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	amount REAL NOT NULL,
	purchase_date INTEGER NOT NULL,
	purchase_price INTEGER NOT NULL);
	`
	_, err := repo.Conn.Exec(query)
	return err
}

func (repo *SQLiteRepository) InsertHoldings(h Holdings) (*Holdings, error) {
	stmt := `INSERT INTO holdings (amount, purchase_date, purchase_price) VALUES (?, ?, ?)`
	result, err := repo.Conn.Exec(stmt, h.Amount, h.PurchaseDate.Unix(), h.PurchasePrice)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errInsertFailed
	}

	h.ID = id
	return &h, nil
}

func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	stmt := `SELECT id, amount, purchase_date, purchase_price FROM holdings`
	rows, err := repo.Conn.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var holdings []Holdings
	var unixTime int64
	for rows.Next() {
		var h Holdings
		err := rows.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
		if err != nil {
			return nil, err
		}
		h.PurchaseDate = time.Unix(unixTime, 0)
		holdings = append(holdings, h)
	}
	return holdings, nil
}

func (repo *SQLiteRepository) GetHoldingByID(id int64) (*Holdings, error) {
	stmt := `SELECT id, amount, purchase_date, purchase_price FROM holdings WHERE id = ?`
	row := repo.Conn.QueryRow(stmt, id)

	var h Holdings
	var unixTime int64
	err := row.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
	if err != nil {
		return nil, err
	}
	h.PurchaseDate = time.Unix(unixTime, 0)
	return &h, nil
}

func (repo *SQLiteRepository) UpdateHolding(id int64, updated Holdings) error {
	if id == 0 {
		return errors.New("invalid id")
	}

	stmt := `UPDATE holdings SET amount = ?, purchase_date = ?, purchase_price = ? WHERE id = ?`
	res, err := repo.Conn.Exec(stmt, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

func (repo *SQLiteRepository) DeleteHolding(id int64) error {
	stmt := `DELETE FROM holdings WHERE id = ?`
	res, err := repo.Conn.Exec(stmt, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}
	return nil
}
