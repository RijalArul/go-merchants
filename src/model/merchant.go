package model

type Merchant struct {
	ID            string `json:"id"`             // Unique ID merchant
	Name          string `json:"name"`           // Nama merchant
	AccountNumber string `json:"account_number"` // Nomor akun merchant
}