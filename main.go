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

	log.Printf("listening on port %d\n", p.port)
	http.ListenAndServe(":"+strconv.Itoa(p.port), mux)
}
