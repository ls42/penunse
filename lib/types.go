package penunse

// User represents a user of this software
type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	First string `json:"first"`
	Pass  string
}

// Transaction is an action that affects your depot
type Transaction struct {
	ID     string   `json:"json"`
	Amount float32  `json:"amount"`
	Tags   []string `json:"tags"`
	Note   string   `json:"note"`
}
