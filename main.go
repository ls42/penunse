package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// User represents a user of this software
type User struct {
	gorm.Model
	Login string `gorm:"unique"`
	First string
	Pass  string
}

// Transaction is an action that affects your depot
type Transaction struct {
	gorm.Model
	Amount float32
	Tags   []string
	Note   string
}

var dbLocation string = "data/"

func main() {
	_ = os.Mkdir(dbLocation)
	db, err := gorm.Open("sqlite3", dbLocation+"penunse.db")
	if err != nil {
		log.Fatal("cant' connect to database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	// db.Create(&User{Login: "kerstin", First: "Kerstin"})
	// db.Create(&User{Login: "stephan", First: "Stephan"})

	// Read
	var user User
	fmt.Printf("%+v", db.First(&user, 1)) // Find user with id 1

}
