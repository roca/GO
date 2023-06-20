package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type Token struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	FirstName string    `db:"first_name" json:"first_name"`
	Email     string    `db:"email" json:"email"`
	PlainText string    `db:"token" json:"token"`
	Hash      []byte    `db:"token_hash" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Expires   time.Time `db:"expiry" json:"expiry"`
}

func (t *Token) Table() string {
	return "tokens"
}

func (t *Token) GetUserForToken(token string) (*User, error) {
	var u User
	var theToken Token

	collection := upper.Collection(t.Table())
	err := collection.
		Find(up.Cond{"token =": token}).
		One(&theToken)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	collection = upper.Collection(u.Table())
	err = collection.
		Find(up.Cond{"id =": theToken.UserID}).
		One(&u)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	u.Token = theToken

	return &u, nil
}

func (t *Token) GetTokensForUser(id int) ([]*Token, error) {
	var all []*Token

	collection := upper.Collection(t.Table())
	err := collection.
		Find(up.Cond{"user_id =": id}).
		All(&all)
	if err != nil {
		return nil, err
	}

	return all, nil
}

// Get token by ID
func (t *Token) Get(id int) (*Token, error) {
	var theToken Token

	collection := upper.Collection(t.Table())
	err := collection.
		Find(up.Cond{"id =": id}).
		One(&theToken)
	if err != nil {
		return nil, err
	}

	return &theToken, nil
}

// Get token by token string
func (t *Token) GetByToken(token string) (*Token, error) {
	var theToken Token

	collection := upper.Collection(t.Table())
	err := collection.
		Find(up.Cond{"token =": token}).
		One(&theToken)
	if err != nil {
		return nil, err
	}

	return &theToken, nil
}

// Delete token by ID
func (t *Token) Delete(id int) error {
	collection := upper.Collection(t.Table())
	err := collection.
		Find(up.Cond{"id =": id}).
		Delete()
	if err != nil {
		return err
	}

	return nil
}

// Delete token by token string
func (t *Token) DeleteByToken(token string) error {
	collection := upper.Collection(t.Table())
	err := collection.
		Find(up.Cond{"token =": token}).
		Delete()
	if err != nil {
		return err
	}

	return nil
}

// Insert token
func (t *Token) Insert(token Token, u User) error {
	collection := upper.Collection(t.Table())

	// delete existing tokens for this user
	err := collection.
		Find(up.Cond{"user_id =": u.ID}).
		Delete()
	if err != nil {
		return err
	}

	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	token.FirstName = u.FirstName
	token.Email = u.Email
	token.UserID = u.ID

	// insert new token
	_, err = collection.
		Insert(&token)
	if err != nil {
		return err
	}

	return nil
}
