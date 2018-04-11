package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 4202, "port to listen on")

	db := &DB{}
	if err := db.Open("penunse.bolt", 0600); err != nil {
		log.Fatal("cannot connect to database `penunse.bolt`")
	}
	defer db.Close()

	mux := http.NewServeMux()

	// TODO: Route to login or main view here
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/api/transaction/read/all", makeHandler(apiAllTransactions, db))
	mux.HandleFunc("/api/transaction/create", makeHandler(apiInsertTransaction, db))

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))),
	)
	log.Printf("Listening on port %d\n", *port)
	http.ListenAndServe(":"+strconv.Itoa(*port), mux)
}
