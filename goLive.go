package goLive

import (
	// "encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Boofny/goLive/logging"
)

//here and here trying to make types methods

// type FunctionHandler func(w http.ResponseWriter, r *http.Request)error //custom handler defined for error handling
// type FunctionHandler func(c *Context)error //custom handler defined for error handling

type FunctionHandler func(c *Context)error //custom handler defined for error handling
type GoLive struct{
	Mux *http.ServeMux
}

func New()*GoLive{ 
	return &GoLive{
		Mux: http.NewServeMux(),
	}
}

func (g *GoLive) GET(path string, /*mux *http.ServeMux,*/ handle FunctionHandler) {
  g.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }

		ctx := &Context{
			Writer: w,
			Request: r,
		}
		err := handle(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
  })
}
func (g *GoLive)POST(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){
  g.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost{
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }

		ctx := &Context{
			Writer: w,
			Request: r,
		}
		err := handle(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
  })
}
func (g *GoLive)DELETE(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){
  g.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete{
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }

		ctx := &Context{
			Writer: w,
			Request: r,
		}
		err := handle(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
  })
}
func (g *GoLive)PUT(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){
  g.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut{
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }

		ctx := &Context{
			Writer: w,
			Request: r,
		}
		err := handle(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
  })
}

// func (g *GoLive)SendJSON(w http.ResponseWriter, status int, data any)error { 
// 	//made this public change it back if you want the scope of the code 
// 	//to be in this file alone not in the main
//   w.Header().Set("Content-Type", "application/json")
//   w.WriteHeader(status)
// 	return json.NewEncoder(w).Encode(data)
// }
//
// func (c *Context) JSON(w http.ResponseWriter, status int, data any) error {
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(status)
//     return json.NewEncoder(w).Encode(data)
// }
//
// func (c *Context) STRING(w http.ResponseWriter, status int, data string)error{
//     w.Header().Set("Content-Type", "text/plain")
// 	w.WriteHeader(status)
// 	_, err := w.Write([]byte(data))
// 	return err 
// }

//needed function

// type GoLive struct{
// 	Mux *http.ServeMux
// }

func (g *GoLive)StartServer(port string, /*mux *http.ServeMux*/) {
	server := &http.Server{
		Addr:    port,
		Handler: logging.Logging(g.Mux), //this is where the output for Requests are
	}

	icon :=  `
 ██████╗  ██████╗ ██╗     ██╗██╗   ██╗███████╗██╗
██╔════╝ ██╔═══██╗██║     ██║██║   ██║██╔════╝██║
██║  ███╗██║   ██║██║     ██║██║   ██║█████╗  ██║
██║   ██║██║   ██║██║     ██║╚██╗ ██╔╝██╔══╝  ╚═╝
╚██████╔╝╚██████╔╝███████╗██║ ╚████╔╝ ███████╗██╗
 ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═══╝  ╚══════╝╚═╝ 
	`    
	blue := "\033[34m"
	yellow := "\033[33m"
	reset := "\033[30m"
	// redH:= "\033[41m"
	// greenH := "\033[42m"
	fmt.Println(blue, icon)
	fmt.Print("\033[34m >>> \033[0m")
	fmt.Print("Server started successfully on port" +  yellow + port + reset)
	fmt.Println("\033[34m <<< \033[0m")

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
	} else if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
	}
}
