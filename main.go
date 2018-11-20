package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	p := parseFlags()
	db := prepareDB(&p)
	defer db.Close()

	mux := http.NewServeMux()

	// TODO: Route to login or main view here
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/api/transaction/read", makeHandler(apiAllTransactions, db))
	mux.HandleFunc("/api/transaction/create", makeHandler(apiInsertTransaction, db))

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))),
	)

	log.Printf("Listening on port %d\n", p.port)
	http.ListenAndServe(":"+strconv.Itoa(p.port), mux)
}
