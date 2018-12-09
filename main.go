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
	mux.HandleFunc("/", makeHandler(appHandler, db))
	mux.HandleFunc("/api/transaction/create", makeHandler(apiUpsertTransaction, db))
	mux.HandleFunc("/api/transaction/read", makeHandler(apiAllTransactions, db))
	mux.HandleFunc("/api/transaction/update", makeHandler(apiUpsertTransaction, db))
	mux.HandleFunc("/api/transaction/delete/", makeHandler(apiDeleteTransaction, db))

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))),
	)

	log.Printf("listening on port %d\n", p.port)
	http.ListenAndServe(":"+strconv.Itoa(p.port), mux)
}
