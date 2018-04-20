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
	Tags    []Tag     `json:"tags" gorm:"many2many:transaction_tags;"`
	Note    string    `json:"note"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Deleted time.Time `json:"deleted"`
}

// Tag is basically just a string
type Tag struct {
	gorm.Model
	Name string
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

// Command line options packed together
type params struct {
	port   int
	dbhost string
	dbport int
	dbuser string
	dbname string
	dbpass string
}

func (p *params) validate() error {
	if p.dbpass == "foo" {
		return errors.New("database password not provided")
	}
	return nil
}
