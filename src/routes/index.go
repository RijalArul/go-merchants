package routes

import (
	"net/http"
)

func IndexRoutes() {
	ApiV1Mux := http.NewServeMux()
    http.Handle("/api/v1/", http.StripPrefix("/api/v1", ApiV1Mux))
	// Set up API v1 routes		
	
	CustomerRoutes(ApiV1Mux)
	PaymentRoutes(ApiV1Mux)
}

