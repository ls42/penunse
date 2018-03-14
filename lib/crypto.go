package penunse

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPass takes a string and return a securely storageable []byte
func EncryptPass(p string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("cannot encrypt provided password string")
	}
	return hash
}
