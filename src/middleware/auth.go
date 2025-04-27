package middleware

import (
	"net/http"
	"strings"

	"go-merchants/src/utils"
)

// AuthMiddleware - Middleware untuk validasi JWT
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
            return
        }

        tokenString := parts[1]
        customerID, err := utils.ValidateJWT(tokenString)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Inject customerID ke context supaya handler bisa pakai
        ctx := r.Context()
        ctx = utils.InjectCustomerID(ctx, customerID)
        r = r.WithContext(ctx)

        next(w, r)
    }
}
