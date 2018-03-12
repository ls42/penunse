package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User represents a user of this software
type User struct {
	gorm.Model
	Login string
	First string
	Pass  string
}

func main() {
	db, err := gorm.Open("sqlite3", "data/penunse.db")
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
