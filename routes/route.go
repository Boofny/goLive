package routes

import (
	"GoLive/logging"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

//here and here trying to make types methods
type HandleFunc func(w http.ResponseWriter, r *http.Request)error //custom handler defined for error handling

type GoLive struct{
}

func New()*GoLive{ 
	return &GoLive{}
}
func (g *GoLive)MuxHandle()http.ServeMux{
	return *http.NewServeMux()
}

//need to make a function that takes the path http function and db connection in order to validate the db functions
// func (g *GoLive)GET(passedFunction func(w http.ResponseWriter, r *http.Request)) {
//
// }

// func (g *GoLive)GET(path string, handle http.HandlerFunc){
// 	http.HandleFunc(path, handle)
// }

func (g *GoLive) GET(path string, mux *http.ServeMux, handle HandleFunc) {
  mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }

		err := handle(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
  })
}
func (g *GoLive)POST(path string, mux *http.ServeMux, handle HandleFunc){
  mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost{
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }

		err := handle(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
  })
}
// func (g *GoLive)GetJsonDefault(w http.ResponseWriter, r *http.Request){
// 	code := map[string]string{
// 		"code": "Hello world",
// 	}
// 	SendJSON(w, http.StatusOK, code)
// }

// func (g *GoLive)PostJsonDefault(w http.ResponseWriter, r *http.Request){
// 	type User struct{
// 		Name string `json:"name"`
// 		Age int `json:"age"`
// 	}
// 	var user User
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		http.Error(w, "Error", http.StatusBadRequest)
// 		return 
// 	}
// 	code := map[string]any {"resp": user.Name, "resp2": user.Age}
// 	SendJSON(w, http.StatusOK, code)
// }

//needed function
func (g *GoLive)SendJSON(w http.ResponseWriter, status int, data any)error { 
	//made this public change it back if you want the scope of the code 
	//to be in this file alone not in the main
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}


//needed function
func (g *GoLive)StartServer(port string, mux *http.ServeMux) {
	//for when needed output
  fmt.Println("\033[31mThis is red text.\033[0m")
	fmt.Println("\033[32mThis is green text.\033[0m")
	fmt.Println("\033[1;34mThis is bold blue text.\033[0m")
	server := &http.Server{
		Addr:    port,
		Handler: logging.Logging(mux), //this is where the output for Requests are
	}

	fmt.Println("Server on port" + port)

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
	} else if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
	}
}
