package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// User represents a user of this software
type User struct {
	ID        int        `json:"id" gorm:"primary_key"`
	Login     string     `json:"login"`
	First     string     `json:"first"`
	Pass      []byte     `json:"pass"`
	CreatedAt time.Time  `json:"created" gorm:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updated" gorm:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:"deleted"`
}

// Transaction is an action that affects your depot
type Transaction struct {
	ID        int        `json:"id" gorm:"primary_key"`
	User      int        `json:"user_id"`
	Amount    float32    `json:"amount"`
	Tags      []Tag      `json:"tags" gorm:"many2many:transactions_tags;"`
	Note      string     `json:"note"`
	CreatedAt time.Time  `json:"created" gorm:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updated" gorm:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:"deleted"`
}

// Tag is basically just a string
type Tag struct {
	ID        int        `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created" gorm:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updated" gorm:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:"deleted"`
}

// Create saves this Transaction to the database
func (t *Transaction) Create(db *gorm.DB) {
	db.Create(t)
}

// HumanDate returns a human readable date
func (t *Transaction) HumanDate(db *gorm.DB) string {
	return t.CreatedAt.Format("02.01.2006")
}

// Username returns the corresponding username for an ID
func (t *Transaction) UserName(db *gorm.DB) string {
	if t.User == 0 {
		return "Stephan"
	} else {
		return "Kerstin"
	}
}

// AmountNormal normalizes the amount to two digits.
func (t *Transaction) AmountNormal(db *gorm.DB) string {
	return fmt.Sprintf("%8.2f", t.Amount)
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

// Check if all params are useful
func (p *params) validate() error {
	// Port must be an int and between 1 and 65535
	validPort := uint16(p.port)
	if validPort < 1 {
		return errors.New("Port invalid, choose one from 1 to 65535")
	}
	return nil
}
