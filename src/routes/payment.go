package routes

import (
	"go-merchants/src/controller"
	"go-merchants/src/middleware"
	"net/http"
)

func PaymentRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/payment", middleware.AuthMiddleware(controller.PaymentHandler))
}
