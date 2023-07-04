package data

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"net/http"
	"strings"
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
		if err != up.ErrNilRecord || err != up.ErrNoMoreRows {
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

// Generate a token
func (t *Token) GenerateToken(userID int, ttl time.Duration) (*Token, error) {
	token := Token{
		UserID:  userID,
		Expires: time.Now().Add(ttl),
	}

	randaomBytes := make([]byte, 16)
	_, err := rand.Read(randaomBytes)
	if err != nil {
		return nil, err
	}

	token.PlainText = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randaomBytes)
	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]

	return &token, nil
}

// AuthenticateToken takes a request and returns a user if the token is valid
// or an error if it is not
func (t *Token) AuthenticateToken(r *http.Request) (*User, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return nil, errors.New("Authorization header required")
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("Authorization header required")
	}

	tokenString := headerParts[1]

	if len(tokenString) != 26 {
		return nil, errors.New("Authorization header required")
	}

	tkn, err := t.GetByToken(tokenString)
	if err != nil {
		return nil, errors.New("No matching token found")
	}

	if tkn.Expires.Before(time.Now()) {
		return nil, errors.New("Token expired")
	}

	user, err := t.GetUserForToken(tokenString)
	if err != nil {
		return nil, errors.New("No matching user found")
	}

	return user, nil
}

// ValidToken

func (t *Token) ValidToken(token string) (bool, error) {
	user, err := t.GetUserForToken(token)
	if err != nil {
		return false, errors.New("No matching user found")
	}

	if user.Token.PlainText == "" {
		return false, errors.New("No matching token found")
	}

	if user.Token.Expires.Before(time.Now()) {
		return false, errors.New("Token expired")
	}

	return true, nil
}
