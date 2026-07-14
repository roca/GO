package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// RememberToken struct
type RememberToken struct {
	ID            int       `db:"id,omitempty"`
	UserID        int       `db:"user_id"`
	RememberToken string    `db:"remember_token"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *RememberToken) Table() string {
	return "remember_tokens"
}

func (t *RememberToken) InsertToken(userID int, token string) error {
	_, err := t.insert(RememberToken{
		UserID:        userID,
		RememberToken: token,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *RememberToken) DeleteToken(rememberToken string) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(up.Cond{"remember_token": rememberToken})
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// GetAll gets all records from the database, using upper
// func (t *RememberToken) getAll(condition up.Cond) ([]*RememberToken, error) {
// 	collection := upper.Collection(t.Table())
// 	var all []*RememberToken

// 	res := collection.Find(condition)
// 	err := res.All(&all)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return all, err
// }

// // Get gets one record from the database, by id, using upper
// func (t *RememberToken) get(id int) (*RememberToken, error) {
// 	var one RememberToken
// 	collection := upper.Collection(t.Table())

// 	res := collection.Find(up.Cond{"id": id})
// 	err := res.One(&one)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &one, nil
// }

// // Update updates a record in the database, using upper
// func (t *RememberToken) update(m RememberToken) error {
// 	m.UpdatedAt = time.Now()
// 	collection := upper.Collection(t.Table())
// 	res := collection.Find(m.ID)
// 	err := res.Update(&m)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Delete deletes a record from the database by id, using upper
// func (t *RememberToken) delete(id int) error {
// 	collection := upper.Collection(t.Table())
// 	res := collection.Find(id)
// 	err := res.Delete()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// Insert inserts a model into the database, using upper
func (t *RememberToken) insert(m RememberToken) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertedID(res.ID())

	return id, nil
}

// Builder is an example of using upper's sql builder
// func (t *RememberToken) builder(id int) ([]*RememberToken, error) {
// 	collection := upper.Collection(t.Table())

// 	var result []*RememberToken

// 	err := collection.Session().
// 		SQL().
// 		SelectFrom(t.Table()).
// 		Where("id > ?", id).
// 		OrderBy("id").
// 		All(&result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
