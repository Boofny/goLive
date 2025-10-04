// TODO: next functions to work on is EnsureAdmin and Authorization for requests
package middleware

import (
	// "log"
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

// type Middleware func(http.Handler) http.Handler //just to remind what middleware defines

func CORS() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

  		if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return // stop here, don’t call next
      }

			next.ServeHTTP(w, r)
		})
	}
}

// TODO: need to add an option for multi origin
//should prob be an array aka slice that contains multiple origins 
//can remake this function but with this in mind

func CustomCORS(allowedOrigin string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

  		if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return // stop here, don’t call next
      }

			next.ServeHTTP(w, r)
		})
	}
}

// func CORSTEST(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func CheckPermissions(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Checking Permissions")
// 		next.ServeHTTP(w, r)
// 	})
// }
