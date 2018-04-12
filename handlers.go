package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request, *DB), db *DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}

// Deliver the reference JavaScript application
func mainHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "main", nil)
}

func apiAllTransactions(w http.ResponseWriter, r *http.Request, db *DB) {
	ts := GetTransactions(db)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(ts)
}

func apiInsertTransaction(w http.ResponseWriter, r *http.Request, db *DB) {
	reqBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	log.Printf("%+v", string(reqBody))
	if err != nil {
		http.Error(w, "cannot read data", 400)
		return
	}
	var t Transaction
	err = json.Unmarshal(reqBody, &t)
	if err != nil {
		http.Error(w, "invalid json", 400)
		return
	}
	err = t.Save(db)
	if err != nil {
		http.Error(w, "cannot save entry", 500)
		return
	}
}
func apiDeleteTransaction(w http.ResponseWriter, r *http.Request, db *DB) {
	userData := r.URL.Path[len("/api/transaction/delete/"):]
	log.Printf("been asked to remove ID `%s`", userData)
	if err := DeleteTransaction([]byte(userData), db); err != nil {
		errorMsg := fmt.Sprintf("transaction `%s` doesn't exist", userData)
		http.Error(w, errorMsg, 400)
	}
}
