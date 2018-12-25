package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	case "/":
		filter :=
			"date(date) >= date('now', 'start of month') AND " +
				"date(date) <= date('now', '+1 month', 'start of month', '-1 day')"
		ts := GetTransactionsWithFilters(db, filter)
		renderTemplate(w, "app", &ts)
	default:
		http.NotFound(w, r)
	}
}

func filterHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	filterString := r.FormValue("timeFilter")
	// TODO: Create new function that takes a string and returns a filter string that
	//       can be passed to GetTransactionsWithFilters
	log.Println(filterString)
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

// Delete an entry
func deleteHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	id, err := strconv.Atoi(r.URL.Path[len("/delete/"):])
	if err != nil {
		log.Printf("%s\n", err)
		http.NotFound(w, r)
	} else {
		log.Printf("deleting entry `%d`\n", id)
		t := GetTransaction(id, db)
		t.Delete(db)
		http.Redirect(w, r, "/", http.StatusSeeOther)
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
		t.Date, err = time.Parse("02.01.2006", r.FormValue("t_date"))
		if err != nil {
			log.Printf("%s\n", err)
			renderTemplate(w, "edit-form", &t)
		}
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
