package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
	"github.com/jinzhu/gorm"
)

// DB is a clone of the standard bolt.DB
type DB struct {
	*bolt.DB
}

// Open initializes the database
func (db *DB) Open(path string, mode os.FileMode) error {
	var err error
	db.DB, err = bolt.Open(path, mode, nil)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("transactions")); err != nil {
			return errors.New("error creating bucket `transactions")
		}
		if _, err = tx.CreateBucketIfNotExists([]byte("users")); err != nil {
			return errors.New("error creating bucket `users")
		}
		return nil
	})
	if err != nil {
		db.Close()
		return err
	}
	return nil
}

// GetTransactions loads all transactions from the database
func GetTransactions(db *gorm.DB) []Transaction {
	var ts []Transaction
	db.Find(&ts)
	for _, trans := range ts {
		fmt.Printf("%+v\n", trans)
	}
	return ts
}

// GetTransaction loads a single transaction from the database, by ID.
func GetTransaction(id int, db *gorm.DB) Transaction {
	var t Transaction
	db.Find(&t, id)
	return t
}

// DeleteTransaction deletes a transaction from the database, by ID.
func DeleteTransaction(id []byte, db *DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("transactions")).Delete(id)
	})
}
