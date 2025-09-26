package middleware

import(
	"net/http"
	"log"
)

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Enabling CORS")
		next.ServeHTTP(w, r)
	})
}
