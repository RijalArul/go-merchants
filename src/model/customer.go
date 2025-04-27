package model

type Customer struct {
	ID       string `json:"id"`        // Unique ID customer
	Username string `json:"username"`  // Username untuk login
	Password string `json:"password"`  // Password yang akan di-hash
	Balance  int    `json:"balance"`   // Saldo customer
	IsLogged bool   `json:"is_logged"` // Status login (tracking di memory atau file)
}
