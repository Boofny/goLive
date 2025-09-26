package middleware

import (
	"net/http"
)


type CORSConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
}

// DefaultCORSConfig returns default CORS configuration
func DefaultCORSConfig() CORSConfig {
	return CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		AllowCredentials: false,
	}
}
type HTTPContext interface {
    GetResponseWriter() http.ResponseWriter
    GetRequest() *http.Request
}

// Your middleware function type
type MiddlewareFunc func(http.Handler) http.Handler

// CORS middleware
func CORS() MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			
			if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusNoContent)
					return
			}
			
			next.ServeHTTP(w, r)
		})
	}
}
