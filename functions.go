package main

import (
	"encoding/binary"
	"flag"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
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

func parseTags(tagString string, t *Transaction) error {
	tags := strings.Split(tagString, ",")
	t.Tags = t.Tags[:0]
	for _, tag := range tags {
		var newTag Tag
		newTag.Name = strings.TrimSpace(tag)
		newTag.Name = strings.ToLower(newTag.Name)
		t.Tags = append(t.Tags, newTag)
	}
	return nil
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
