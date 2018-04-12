package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/boltdb/bolt"
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
func GetTransactions(db *DB) []Transaction {
	var ts []Transaction
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		c := b.Cursor()
		bucketContainsData, _ := c.First()
		if bucketContainsData == nil {
			ts = append(ts, Transaction{})
			return nil
		}
		b.ForEach(func(key, value []byte) error {
			var t Transaction
			err := json.Unmarshal(value, &t)
			if err != nil {
				log.Fatal(err)
			}
			ts = append(ts, t)
			return nil
		})
		return nil
	})
	return ts
}

// GetTransaction loads a single transaction from the database, by ID.
func GetTransaction(id string, db *DB) Transaction {
	var t Transaction
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		value := b.Get([]byte(id))
		err := json.Unmarshal(value, &t)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	return t
}

// DeleteTransaction deletes a transaction from the database, by ID
func DeleteTransaction(id []byte, db *DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("transactions")).Delete(id)
	})
	// TODO
	// Looks like the id does not really match the id from the database
	// debug here why that's the case -> id should be exactly the id of
	// transaction
}
