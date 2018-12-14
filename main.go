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
	mux.HandleFunc("/", makeHandler(mainHandler, db))
	mux.HandleFunc("/add", makeHandler(addHandler, db))
	mux.HandleFunc("/edit/", makeHandler(editHandler, db))
	mux.HandleFunc("/save/", makeHandler(saveHandler, db))
	mux.HandleFunc("/delete/", makeHandler(deleteHandler, db))
	mux.HandleFunc("/api/transaction/create", makeHandler(apiUpsertTransaction, db))
	mux.HandleFunc("/api/transaction/read", makeHandler(apiAllTransactions, db))
	mux.HandleFunc("/api/transaction/update", makeHandler(apiUpsertTransaction, db))
	mux.HandleFunc("/api/transaction/delete/", makeHandler(apiDeleteTransaction, db))

	log.Printf("listening on port %d\n", p.port)
	http.ListenAndServe(":"+strconv.Itoa(p.port), mux)
}
