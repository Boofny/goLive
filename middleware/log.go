package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
type wrappedWrite struct{
	http.ResponseWriter
	satusCode int
	wroteHeader bool
}

func (w *wrappedWrite)WriteHeader(statusCode int){ //need to look into this 
	if w.wroteHeader {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.satusCode = statusCode
	w.wroteHeader = true
}//is this recursive??????

func (w *wrappedWrite) Write(b []byte) (int, error) {
  if !w.wroteHeader {
    w.WriteHeader(http.StatusOK)
  }
  return w.ResponseWriter.Write(b)
}

//NOTE: commenting this out in order to test something if want to go back to optional logger uncommnt this and see the golive file
// func Logger() Middleware{
// 	return func (next http.Handler)http.Handler  { //og name is Logging() new name Logger()
// 		redH:= "\033[31m"
// 		greenH := "\033[32m"
// 		reset := "\033[0m"
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 			start := time.Now()
//
// 			wrapped := &wrappedWrite{
// 				ResponseWriter: w,
// 				satusCode: http.StatusOK,
// 			}
// 			next.ServeHTTP(wrapped, r)
// 			code := wrapped.satusCode
// 			if code >= 400 && code <= 599{ 
// 				fmt.Print("\033[31m >>> \033[0m") //error
// 				log.Println(redH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start))
// 			}else{
// 				fmt.Print("\033[32m >>> \033[0m") //good
// 				log.Println(greenH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start)) 
// 			}
// 		})
// 	}
// }

func Logger(next http.Handler)http.Handler  { //og name is Logging() new name Logger()
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
		code := wrapped.satusCode
		if code >= 400 && code <= 599{ 
			fmt.Print("\033[31m >>> \033[0m") //error
			log.Println(redH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start))
		}else{
			fmt.Print("\033[32m >>> \033[0m") //good
			log.Println(greenH, wrapped.satusCode, reset, r.Method , r.URL.Path, time.Since(start)) 
		}
	})
}




