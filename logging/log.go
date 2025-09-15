package logging

import (
	"fmt"
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
	redH:= "\033[31m"
	greenH := "\033[32m"
	reset := "\033[0m"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()

		wrapped := &wrappedWrite{
			ResponseWriter: w,
			satusCode: http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		if wrapped.satusCode != 200{
			fmt.Print("\033[31m >>> \033[0m")
			log.Println(redH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start))
		}else{
			fmt.Print("\033[32m >>> \033[0m")
			log.Println(greenH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start))
		}
	})
}
