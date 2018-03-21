package penunse

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
)

// GetAllTransactions from the database
func GetAllTransactions(db *bolt.DB) []Transaction {
	var transactions []Transaction
	var transaction Transaction
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		b.ForEach(func(key, value []byte) error {
			err := json.Unmarshal(value, &transaction)
			if err != nil {
				log.Fatal(err)
			}
			transactions = append(transactions, transaction)
			return nil
		})
		return nil
	})
	return transactions
}

// GetTransaction from the database
func GetTransaction(id string, db *bolt.DB) Transaction {
	var transaction Transaction
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		value := b.Get([]byte(id))
		err := json.Unmarshal(value, &transaction)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	return transaction
}

// SaveTransaction to the database
func SaveTransaction(transaction Transaction, db *bolt.DB) {
	id := transaction.ID
	encodedTransaction, err := json.Marshal(transaction)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		err := b.Put([]byte(id), encodedTransaction)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
}

// DeleteTransaction from the database
func DeleteTransaction(id string, db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("transactions"))
		err := b.Delete([]byte(id))
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
}
