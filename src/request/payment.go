package request

type PaymentRequest struct {
	MerchantID string `json:"merchant_id"` // ID merchant tujuan
	Amount     int    `json:"amount"`      // Jumlah uang yang dibayar
}
