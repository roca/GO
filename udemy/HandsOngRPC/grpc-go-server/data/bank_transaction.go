package data

import (
	"time"

	"github.com/google/uuid"
	up "github.com/upper/db/v4"
)

// BankTransaction struct
type BankTransaction struct {
	ID        uuid.UUID `db:"transaction_uuid,omitempty"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *BankTransaction) Table() string {
	return "bank_transactions"
}

// GetAll gets all records from the database, using upper
func (t *BankTransaction) GetAll(condition up.Cond) ([]*BankTransaction, error) {
	collection := upper.Collection(t.Table())
	var all []*BankTransaction

	res := collection.Find(condition)
	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *BankTransaction) Get(id int) (*BankTransaction, error) {
	var one BankTransaction
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *BankTransaction) Update(m BankTransaction) error {
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *BankTransaction) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *BankTransaction) Insert(m BankTransaction) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
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
func (t *BankTransaction) Builder(id int) ([]*BankTransaction, error) {
	collection := upper.Collection(t.Table())

	var result []*BankTransaction

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
