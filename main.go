package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/password"
	"github.com/qor/session/manager"
)

func main() {
	p := parseFlags()
	db := prepareDB(&p)
	defer db.Close()

	// Initialize Auth with configuration
	Auth := auth.New(&auth.Config{
		DB: db,
	})

	// Migrate AuthIdentity model, AuthIdentity will be used to save auth info, like username/password, oauth token, you could change thatgormDB.AutoMigrate(&auth_identity.AuthIdentity{})
	db.AutoMigrate(&auth_identity.AuthIdentity{})
	// Register Auth providers
	// Allow use username/password
	Auth.RegisterProvider(password.New(&password.Config{}))

	var firstEntry Transaction
	db.First(&firstEntry, 1)
	log.Printf("%+v\n", firstEntry)

	mux := http.NewServeMux()

	// TODO: Route to login or main view here
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/api/transaction/read", makeHandler(apiAllTransactions, db))
	// mux.HandleFunc("/api/transaction/read/", makeHandler(apiTransaction, db))
	// mux.HandleFunc("/api/transaction/create", makeHandler(apiInsertTransaction, db))
	// mux.HandleFunc("/api/transaction/delete/", makeHandler(apiDeleteTransaction, db))

	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))),
	)
	mux.Handle("/auth/", Auth.NewServeMux())

	log.Printf("Listening on port %d\n", p.port)
	http.ListenAndServe(":"+strconv.Itoa(p.port), mux)
	http.ListenAndServe(":9000", manager.SessionManager.Middleware(mux))
}
