package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

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
		renderTemplate(w, "app", &ts)
	default:
		http.NotFound(w, r)
	}
}

// Render the add-form
func addHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	log.Println("creating empty transaction for `add`-form")
	id := 0
	t := GetTransaction(id, db)
	renderTemplate(w, "edit-form", &t)
}

// Render the edit-form
func editHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id, err := strconv.Atoi(r.URL.Path[len("/edit/"):])
	if err != nil {
		http.NotFound(w, r)
	} else {
		log.Printf("requested to edit entry `%d`\n", id)
		t := GetTransaction(id, db)
		renderTemplate(w, "edit-form", &t)
	}

}

// Save entry and redirect to mainHandler
func saveHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id, err := strconv.Atoi(r.URL.Path[len("/save/"):])
	if err != nil {
		log.Printf("%s\n", err)
		http.NotFound(w, r)
	} else {
		t := GetTransaction(id, db)
		t.User, err = strconv.Atoi(r.FormValue("t_user"))
		if err != nil {
			log.Printf("%s\n", err)
			renderTemplate(w, "edit-form", &t)
		}
		t.Note = r.FormValue("t_note")
		amount64, err := strconv.ParseFloat(r.FormValue("t_amount"), 32)
		if err != nil {
			log.Printf("%s\n", err)
			renderTemplate(w, "edit-form", &t)
		}
		t.Amount = float32(amount64)
		tagString := r.FormValue("t_tags")
		tags := strings.Split(tagString, ",")
		db.Model(&t).Association("Tags").Delete(t.Tags)
		t.Tags = t.Tags[:0]
		for _, tag := range tags {
			var newTag Tag
			newTag.Name = strings.TrimSpace(tag)
			t.Tags = append(t.Tags, newTag)
		}
		t.Save(db)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
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
