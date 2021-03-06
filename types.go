package main

import (
	"errors"
	"fmt"
	"strings"
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
	Date      time.Time  `json:"date"`
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

// Create inserts this Transaction to the database
func (t *Transaction) Create(db *gorm.DB) {
	db.Create(&t)
}

// Save saves this Transaction to the database
func (t *Transaction) Save(db *gorm.DB) {
	db.Save(&t)
}

// Delete removes a Transaction from the database
func (t *Transaction) Delete(db *gorm.DB) {
	db.Delete(&t)
}

// HumanDate returns a human readable date
func (t *Transaction) HumanDate() string {
	if t.Date.IsZero() {
		return t.CreatedAt.Format("02.01.2006")
	}
	return t.Date.Format("02.01.2006")
}

// UserName returns the corresponding username for an ID
func (t *Transaction) UserName() string {
	if t.User == 0 {
		return "Stephan"
	}
	return "Kerstin"
}

// AmountNormal normalizes the amount to two digits.
func (t *Transaction) AmountNormal() string {
	return fmt.Sprintf("%.2f", t.Amount)
}

// TagsNormal combines all tags to a single, comma-separated string
func (t *Transaction) TagsNormal() string {
	var tags []string
	for _, tag := range t.Tags {
		tags = append(tags, tag.Name)
	}
	return strings.Join(tags, ", ")
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
