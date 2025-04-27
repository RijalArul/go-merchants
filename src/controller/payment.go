package controller

import (
	"encoding/json"
	"net/http"

	"go-merchants/src/request"
	"go-merchants/src/service"
	"go-merchants/src/utils"
)

var paymentService = service.NewPaymentService()

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    customerID, ok := utils.GetCustomerID(r.Context())
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var paymentRequest request.PaymentRequest
    err := json.NewDecoder(r.Body).Decode(&paymentRequest)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    errPayment := paymentService.Pay(customerID, paymentRequest.MerchantID, paymentRequest.Amount)
    if errPayment != nil {
        http.Error(w, errPayment.Error(), http.StatusBadRequest)
        return
    }

    response := map[string]interface{}{
        "message": "Payment successful",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
