package main

import (
	"encoding/binary"
	"flag"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// inSlice checks if a string is in a slice of strings
func inSlice(needle string, haystack []string) bool {
	for _, hay := range haystack {
		if needle == hay {
			return true
		}
	}
	return false
}

// parseFlags defines and then parses all command line flags
func parseFlags() params {
	p := params{}
	flag.IntVar(&p.port, "port", 4202, "port to listen on")
	flag.Parse()
	if err := p.validate(); err != nil {
		log.Fatal(err)
	}
	return p
}

func prepareDB(p *params) *gorm.DB {
	db, err := gorm.Open("sqlite3", "./data/penunse.db")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Transaction{})
	return db
}
