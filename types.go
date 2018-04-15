package main

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jinzhu/gorm"
)

// User represents a user of this software
type User struct {
	gorm.Model
	ID      int    `json:"id"`
	Login   string `json:"login"`
	First   string `json:"first"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
	Pass    []byte `json:"pass"`
}

// Transaction is an action that affects your depot
type Transaction struct {
	gorm.Model
	ID      int       `json:"id"`
	User    int       `json:"user_id"`
	Amount  float32   `json:"amount"`
	Tags    []string  `json:"tags"`
	Note    string    `json:"note"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted time.Time `json:"deleted"`
}

// Save saves this Transaction to the database
func (t *Transaction) Save(db *DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		id, _ := b.NextSequence()
		t.ID = int(id)
		t.Created = time.Now()
		t.Updated = time.Now()
		tj, err := json.Marshal(t)
		if err != nil {
			return errors.New("unable to serialize transaction to json")
		}
		err = b.Put(itob(t.ID), tj)
		if err != nil {
			return errors.New("unable to save this transaction")
		}
		return nil
	})
}
