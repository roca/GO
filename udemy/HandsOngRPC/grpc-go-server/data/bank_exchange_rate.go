package data

import (
	"time"

	"github.com/google/uuid"
	up "github.com/upper/db/v4"
)

// BankExchangeRate struct
type BankExchangeRate struct {
	ID                 uuid.UUID `db:"exchange_rate_uuid,omitempty"`
	FromCurrency       string    `db:"from_currency"`
	ToCurrency         string    `db:"to_currency"`
	Rate               float64   `db:"rate"`
	ValidFromTimestamp time.Time `db:" valid_from_timestamp"`
	ValidToTimestamp   time.Time `db:" valid_to_timestamp"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *BankExchangeRate) Table() string {
	return "bank_exchange_rates"
}

// GetAll gets all records from the database, using upper
func (t *BankExchangeRate) GetAll(condition up.Cond) ([]*BankExchangeRate, error) {
	collection := upper.Collection(t.Table())
	var all []*BankExchangeRate

	res := collection.Find(condition)
	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *BankExchangeRate) Get(id int) (*BankExchangeRate, error) {
	var one BankExchangeRate
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *BankExchangeRate) Update(m BankExchangeRate) error {
	// m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *BankExchangeRate) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *BankExchangeRate) Insert(m BankExchangeRate) (int, error) {
	// m.CreatedAt = time.Now()
	// m.UpdatedAt = time.Now()
	m.ID = uuid.New()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertedID(res.ID())

	return id, nil
}

// Builder is an example of using upper's sql builder
func (t *BankExchangeRate) Builder(id int) ([]*BankExchangeRate, error) {
	collection := upper.Collection(t.Table())

	var result []*BankExchangeRate

	err := collection.Session().
		SQL().
		SelectFrom(t.Table()).
		Where("id > ?", id).
		OrderBy("id").
		All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
