package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

//
// type CORSConfig struct {
// 	AllowOrigins     []string
// 	AllowMethods     []string
// 	AllowHeaders     []string
// 	AllowCredentials bool
// }
//
// // DefaultCORSConfig returns default CORS configuration
// func DefaultCORSConfig() CORSConfig {
// 	return CORSConfig{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"},
// 		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
// 		AllowCredentials: false,
// 	}
// }
// type HTTPContext interface {
//     GetResponseWriter() http.ResponseWriter
//     GetRequest() *http.Request
// }
//
// // Your middleware function type
// type MiddlewareFunc func(http.Handler) http.Handler
//
// // CORS middleware
// func CORS() MiddlewareFunc {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.Header().Set("Access-Control-Allow-Origin", "*")
// 			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//
// 			if r.Method == "OPTIONS" {
// 					w.WriteHeader(http.StatusNoContent)
// 					return
// 			}
//
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

//dont know how wise this is but just to define types would prob be fine

// func CORS()string{
// 	return "*"
// }
// func CORS(next http.Handler) http.Handler {
//   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//     // Wrap into your Context
//     c := &Context{Writer: w, Request: r}
//
//     // Set headers
//     c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//     c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
//     c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
//
//     // Handle preflight requests directly
//     if c.Request.Method == http.MethodOptions {
//       c.Writer.WriteHeader(http.StatusNoContent)
//       return
//     }
//
//     // Continue down the chain
//     next.ServeHTTP(c.Writer, c.Request)
//   })
// }
