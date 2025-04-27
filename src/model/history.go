package model

type History struct {
	ID         string `json:"id"`                    // Unique ID untuk setiap event
	CustomerID string `json:"customer_id"`           // Siapa yang melakukan event
	Action     string `json:"action"`                // login, payment, logout
	MerchantID string `json:"merchant_id,omitempty"` // Optional: untuk payment
	Amount     int    `json:"amount,omitempty"`      // Optional: jumlah pembayaran
	Timestamp  string `json:"timestamp"`             // Waktu event
	Success    bool   `json:"success"`               // Status sukses atau gagal
}
