package penunse

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
)

// Users is a collection of User
type Users struct {
	Users []User
}

// User represents a user of this software
type User struct {
	ID      string `json:"id"`
	Login   string `json:"login"`
	First   string `json:"first"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
	Pass    []byte
}

// Transaction is an action that affects your depot
type Transaction struct {
	ID     string   `json:"json"`
	Amount float32  `json:"amount"`
	Tags   []string `json:"tags"`
	Note   string   `json:"note"`
}

// Save saves this Transaction to the database
func (t *Transaction) Save(db *bolt.DB) error {
	tj, err := json.Marshal(t)
	if err != nil {
		return errors.New("unable to marshal transaction")
	}
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		err := b.Put([]byte(t.ID), tj)
		if err != nil {
			return errors.New("unable to save this transaction")
		}
		return nil
	})
}

// Delete removes this Transaction from the database
func (t *Transaction) Delete(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		err := b.Delete([]byte(t.ID))
		if err != nil {
			return errors.New("unable to delete this transaction")
		}
		return nil
	})
}
