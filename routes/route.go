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
type GoLive struct{}
func New()*GoLive{
	return &GoLive{}
}

func (g *GoLive)GetStringDefault(w http.ResponseWriter, r *http.Request){
	sendJSON(w, http.StatusAccepted, "Hello")
}

func (g *GoLive)GET(path string) error {

	return nil
}
//here again making some thing like echo 
//have and idea make two things and http template for fast project starting and a framework like echo 
func GetJsonDefault(w http.ResponseWriter, r *http.Request){
	code := map[string]string{
		"code": "Hello world",
	}
	sendJSON(w, http.StatusOK, code)
}

func PostJsonDefault(w http.ResponseWriter, r *http.Request){
	type User struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return 
	}
	code := map[string]any {"resp": user.Name, "resp2": user.Age}
	sendJSON(w, http.StatusOK, code)
}

//needed function
func sendJSON(w http.ResponseWriter, status int, data any) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  json.NewEncoder(w).Encode(data)
}


//needed function
func StartServer(port string, mux *http.ServeMux) {
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
