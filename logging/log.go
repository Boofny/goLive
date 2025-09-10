package logging

import (
	"log"
	"net/http"
	"time"
)

type wrappedWrite struct{
	http.ResponseWriter
	satusCode int
}

func (w *wrappedWrite) WriteHeader(statusCode int){
	w.ResponseWriter.WriteHeader(statusCode)
	w.satusCode = statusCode
}

func Logging(next http.Handler)http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()

		wrapped := &wrappedWrite{
			ResponseWriter: w,
			satusCode: http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		log.Println(wrapped.satusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
