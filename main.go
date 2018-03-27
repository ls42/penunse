package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"
	"gitlab.ls42.de/go/penunse/lib"
)

type templateConfig struct {
	AppName string
}

const (
	appName = "Penunse"
)

func main() {

	port := flag.Int("port", 4202, "port to listen on")

	// TODO: Move database init to `lib/database.go`
	db, err := bolt.Open("penunse.bolt", 0600, nil)
	if err != nil {
		log.Fatal("can't connect to database")
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("transactions"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	defer db.Close()

	mux := http.NewServeMux()
	t := template.Must(template.ParseGlob("templates/*"))
	tcfg := templateConfig{AppName: appName}

	// TODO: Route to login or main view here
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = t.ExecuteTemplate(w, "main", tcfg)
		if err != nil {
			log.Fatal("rendering default template failed")
		}
	})
	mux.HandleFunc("/api/transaction/read/all", func(w http.ResponseWriter, r *http.Request) {
		// Debugging only, remove later
		for k, v := range r.Header {
			log.Printf("Header %q\t=> %q\n", k, v)
		}

		ts := penunse.GetTransactions(db)
		if err != nil {
			http.Error(w, "cannot serialize transactions", 500)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(ts)
	})
	mux.HandleFunc("/api/transaction/create", func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "cannot read data", 400)
			return
		}
		var t penunse.Transaction
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
	})

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))),
	)
	pass := "little-horror-shop-is-fancy"
	pHash, err := penunse.EncryptPass(pass)
	if err != nil {
		log.Fatal("error while hashing password")
	}
	log.Printf("%s", pHash)
	log.Printf("Listening on port %d\n", *port)
	http.ListenAndServe(":"+strconv.Itoa(*port), mux)
}
