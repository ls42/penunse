package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request, *gorm.DB), db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}

// Deliver a few standard functions and files or 404 if nothing matched
func mainHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	switch r.URL.RequestURI() {
	case "/favicon.ico":
		http.ServeFile(w, r, "static/favicon.ico")
	case "/":
		ts := GetTransactions(db)
		renderTemplate(w, "app", ts)
	default:
		http.NotFound(w, r)
	}
}

// Render the edit-form
func editHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id, err := strconv.Atoi(r.URL.Path[len("/edit/"):])
	if err != nil {
		http.NotFound(w, r)
	} else {
		log.Printf("requested to edit entry `%d`\n", id)
		t := GetTransaction(id, db)
		renderTemplate(w, "edit-form", t)
	}

}

// Save entry and redirect to mainHandler
func saveHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// 1. parse ID from URL
	// 2. Parse form data
	// 3. Update entry in database
}

func apiAllTransactions(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	ts := GetTransactions(db)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(ts)
}

func apiUpsertTransaction(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	reqBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	log.Printf("new entry:\t%+v", string(reqBody))
	if err != nil {
		http.Error(w, "cannot read data", 400)
		return
	}
	var t Transaction
	err = json.Unmarshal(reqBody, &t)
	if err != nil {
		fmt.Printf("%v", err)
		http.Error(w, "invalid json", 400)
		return
	}
	t.Create(db)
}

func apiDeleteTransaction(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	transactionID, err := strconv.Atoi(r.URL.Path[len("/api/transaction/delete/"):])
	if err != nil {
		http.NotFound(w, r)
	}
	var t Transaction
	if err = db.First(&t, transactionID).Error; err != nil {
		http.NotFound(w, r)
	} else {
		db.Delete(&t)
	}
}
