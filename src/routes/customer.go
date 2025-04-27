package routes

import (
	"go-merchants/src/controller"
	"go-merchants/src/middleware"
	"net/http"
)

func CustomerRoutes(mux *http.ServeMux) {
    // Customer Authentication Routes
	mux.HandleFunc("/login", controller.LoginHandler)
	mux.HandleFunc("/logout", middleware.AuthMiddleware(controller.LogoutHandler))
}
