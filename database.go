package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// GetTransactions loads all transactions from the database
func GetTransactions(db *gorm.DB) []Transaction {
	var ts []Transaction
	db.Table("transactions").
		Preload("Tags").
		Find(&ts)
	return ts
}

// GetTransaction loads a single transaction from the database, by ID.
func GetTransaction(id int, db *gorm.DB) Transaction {
	var t Transaction
	err := db.Table("transactions").
		Preload("Tags").
		First(&t, id).
		Error
	if err != nil {
		t.Date = time.Now()
	}
	return t
}
