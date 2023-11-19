package data

import (
	"time"

	"github.com/google/uuid"
	up "github.com/upper/db/v4"
)

// BankAccount struct
type BankAccount struct {
	ID             uuid.UUID `db:"account_uuid,omitempty"`
	AccountNumber  string    `db:"account_number"`
	AccountName    string    `db:"account_name"`
	Currency       string    `db:"currency"`
	CurrentBalance float64   `db:"current_balance"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *BankAccount) Table() string {
	return "bank_accounts"
}

// GetAll gets all records from the database, using upper
func (t *BankAccount) GetAll(condition up.Cond) ([]*BankAccount, error) {
	collection := upper.Collection(t.Table())
	var all []*BankAccount

	res := collection.Find(condition)
	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *BankAccount) Get(id string) (*BankAccount, error) {
	var one BankAccount
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"account_uuid": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *BankAccount) Update(m BankAccount) error {
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
func (t *BankAccount) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *BankAccount) Insert(m BankAccount) (up.ID, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	m.ID = uuid.New()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	//id := getInsertedID(res.ID())

	return res.ID, nil
}

// Builder is an example of using upper's sql builder
func (t *BankAccount) Builder(id int) ([]*BankAccount, error) {
	collection := upper.Collection(t.Table())

	var result []*BankAccount

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
