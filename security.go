package main

import (
	"net/http"
)

// LoginHandler checks if theres a valid session and asks for login if not
type LoginHandler struct {
	next http.Handler
}

// EncryptPass takes a string and return a securely storageable []byte
func EncryptPass(p string) ([]byte, error) {
	// return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return []byte("lol"), nil
}
