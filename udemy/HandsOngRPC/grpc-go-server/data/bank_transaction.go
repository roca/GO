package data

import (
	"errors"
	"time"

	"github.com/google/uuid"
	up "github.com/upper/db/v4"
)

// BankTransaction struct
type BankTransaction struct {
	ID uuid.UUID `db:"transaction_uuid,omitempty"`

	AccountID            uuid.UUID `db:"account_uuid"`
	TransactionTimestamp time.Time `db:"transaction_timestamp"`
	Amount               float64   `db:"amount"`
	TransactionType      string    `db:"transaction_type"`
	Notes                string    `db:"notes"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
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
func (t *BankTransaction) Insert(m BankTransaction) (up.ID, error) {
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

func (t *BankTransaction) BulkInsert(account BankAccount, m []BankTransaction) (float64, error) {
	balance := account.CurrentBalance
	err := upper.Tx(func(tx up.Session) error {
		for _, v := range m {
			switch v.TransactionType {
			case "DEPOSIT":
				balance += v.Amount
			case "WITHDRAWAL":
				balance -= v.Amount
			}
			if balance < 0 {
				return errors.New("Insufficient balance")
			}
			_, err := t.Insert(v)
			if err != nil {
				return err
			}
		}
		updatedBankAccount := BankAccount{
			ID:             account.ID,
			CurrentBalance: balance,
			AccountNumber:  account.AccountNumber,
			AccountName:    account.AccountName,
			Currency:       account.Currency,
			CreatedAt:      account.CreatedAt,
			UpdatedAt:      time.Now(),
		}
		account.Update(updatedBankAccount)

		return nil
	})
	if err != nil {
		return account.CurrentBalance, err
	}
	return balance, err
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
