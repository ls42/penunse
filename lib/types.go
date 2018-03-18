package penunse

import (
	"time"
)

// Users is a collection of User
type Users struct {
	Users []User
}

// User represents a user of this software
type User struct {
	ID      string `json:"id"`
	Login   string `json:"login"`
	First   string `json:"first"`
	Created time.Time
	Updated time.Time
	Deleted time.Time
	Pass    []byte
}

// Transaction is an action that affects your depot
type Transaction struct {
	ID     string   `json:"json"`
	Amount float32  `json:"amount"`
	Tags   []string `json:"tags"`
	Note   string   `json:"note"`
}
