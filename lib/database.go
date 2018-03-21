package penunse

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
)

// GetTransactions loads all transactions from the database
func GetTransactions(db *bolt.DB) []Transaction {
	var ts []Transaction
	var t Transaction
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		b.ForEach(func(key, value []byte) error {
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
func GetTransaction(id string, db *bolt.DB) Transaction {
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
