package controller

import (
	"encoding/json"
	"go-merchants/src/request"
	"go-merchants/src/response"
	"go-merchants/src/service"
	"go-merchants/src/utils"
	"net/http"
)

var customerService = service.NewCustomerService()

// LoginHandler - Handle login POST
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest request.CustomerLoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := customerService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := response.LoginResponse{
		Token: token,	
		Message: "Login successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// LogoutHandler - Handle logout POST
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    customerID, ok := utils.GetCustomerID(r.Context())
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    err := customerService.Logout(customerID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response := map[string]interface{}{
        "message": "Logout successful",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
