package main

import (
	"html/template"
	"log"
	"net/http"

	"gitlab.ls42.de/go/penunse/lib"

	"github.com/boltdb/bolt"
)

type defaultConfig struct {
	appName string
	port    string
}

var cfg = defaultConfig{
	appName: "Penunse",
	port:    ":4200",
}

func main() {
	db, err := bolt.Open("penunse.bolt", 0600, nil)
	if err != nil {
		log.Fatal("cant' connect to database")
	}
	defer db.Close()

	mux := http.NewServeMux()
	t := template.Must(template.ParseGlob("templates/*"))

	// TODO: Route to login or main view here
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = t.ExecuteTemplate(w, "main", cfg)
		if err != nil {
			log.Fatal("rendering default template failed")
		}
	})

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))),
	)
	pass := "little-horror-shop-is-fancy"
	log.Printf("%s", penunse.EncryptPass(pass))
	log.Printf("Listening on port %s\n", cfg.port)
	http.ListenAndServe(cfg.port, mux)
}
