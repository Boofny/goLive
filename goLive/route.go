package goLive 

import (
	"github.com/Boofny/goLive/logging"
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
//dont find this to needed
// func (g *GoLive)MuxHandle()*http.ServeMux{
// 	return http.NewServeMux()
// }
type Context struct{
	handle *HandleFunc
}
func (c *Context)JSON()error{
	return nil
}

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
func (g *GoLive)DELETE(path string, mux *http.ServeMux, handle HandleFunc){
  mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete{
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
func (g *GoLive)PUT(path string, mux *http.ServeMux, handle HandleFunc){
  mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut{
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
	server := &http.Server{
		Addr:    port,
		Handler: logging.Logging(mux), //this is where the output for Requests are
	}

	fmt.Println( ` 
 ██████╗  ██████╗ ██╗     ██╗██╗   ██╗███████╗██╗
██╔════╝ ██╔═══██╗██║     ██║██║   ██║██╔════╝██║
██║  ███╗██║   ██║██║     ██║██║   ██║█████╗  ██║
██║   ██║██║   ██║██║     ██║╚██╗ ██╔╝██╔══╝  ╚═╝
╚██████╔╝╚██████╔╝███████╗██║ ╚████╔╝ ███████╗██╗
 ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═══╝  ╚══════╝╚═╝ 
	`)
  fmt.Println("\033[1;34mServer started successfully on port" +  server.Addr +"!\033[0m")

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
	} else if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
	}
}
