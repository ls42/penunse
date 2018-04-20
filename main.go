package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	p := parseFlags()

	dbParams := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		p.dbhost,
		p.dbport,
		p.dbuser,
		p.dbname,
		p.dbpass)
	db, err := gorm.Open("postgres", dbParams)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()

	// TODO: Route to login or main view here
	mux.HandleFunc("/", mainHandler)
	// mux.HandleFunc("/api/transaction/read", makeHandler(apiAllTransactions, db))
	// mux.HandleFunc("/api/transaction/read/", makeHandler(apiTransaction, db))
	// mux.HandleFunc("/api/transaction/create", makeHandler(apiInsertTransaction, db))
	// mux.HandleFunc("/api/transaction/delete/", makeHandler(apiDeleteTransaction, db))

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))),
	)
	log.Printf("Listening on port %d\n", p.port)
	http.ListenAndServe(":"+strconv.Itoa(p.port), mux)
}
