package middleware

import (
	"net"
	"net/http"
)

var allowedIPs = []string{
    "127.0.0.1", // Localhost
    "::1",       // Localhost IPv6
    "192.168.1.100", // Misal IP kantor kamu
}

// IPWhitelistMiddleware - Allow only requests from allowed IPs
func IPWhitelistMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip, _, err := net.SplitHostPort(r.RemoteAddr)
        if err != nil {
            http.Error(w, "Invalid IP address", http.StatusForbidden)
            return
        }

        allowed := false
        for _, allowedIP := range allowedIPs {
            if ip == allowedIP {
                allowed = true
                break
            }
        }

        if !allowed {
            http.Error(w, "IP not allowed", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}
