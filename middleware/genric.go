package middleware

import (
	"log"
	"net/http"
	// "strings"
)

// func EnsureAdmin(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Checking if user is admin")
// 		if !strings.Contains(r.Header.Get("Authorization"), "Admin") {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
//
// func LoadUser(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// log.Println("Loading user")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func CORS2 () Middleware{
// 	return func (next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.Header().Set("Access-Control-Allow-Origin", "*")//will need this bro bro
// 			log.Println("Enabling CORS")
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			log.Println("CORS applied")
			next.ServeHTTP(w, r)
		})
	}
}

//TODO need to add an option for multi origin
func CustomCORS(allowedOrigin string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			log.Println("CORS applied")
			next.ServeHTTP(w, r)
		})
	}
}

// func CheckPermissions(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Checking Permissions")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func Default() Middleware {
// 	return CreateStack(
// 		Logging,
// 		AllowCors,
// 	)
// }
