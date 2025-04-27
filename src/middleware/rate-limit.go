package middleware

import (
	"net/http"
	"sync"
	"time"
)

var visitors = make(map[string]time.Time)
var mu sync.Mutex

// RateLimitMiddleware - Allow max 1 request per 5 seconds per IP
func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        lastSeen, exists := visitors[r.RemoteAddr]
        now := time.Now()

        if exists && now.Sub(lastSeen) < 5 *time.Second {
            mu.Unlock()
            http.Error(w, "Too many requests", http.StatusTooManyRequests)
            return
        }

        visitors[r.RemoteAddr] = now
        mu.Unlock()

        next.ServeHTTP(w, r)
    })
}
