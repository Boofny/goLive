package goLive

import (
	// "encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Boofny/goLive/middleware"
)

const (
	GET = http.MethodGet
	POST = http.MethodPost
	PUT = http.MethodPut
	DELETE = http.MethodDelete
)

var (
	ErrInvalidRedirectCode = errors.New("invalid redirect status code") 
	//will add more err code in future as this thing grows
)

type FunctionHandler func(c *Context)error //custom handler defined for error handling

//struct that has a mux handler property
type GoLive struct{
	Mux *http.ServeMux
}

//Method for starting the goLive session
func New()*GoLive{ 
	return &GoLive{
		Mux: http.NewServeMux(),
	}
}

//Use will have to take a paramiter of middeware.CORS() and that should return something http i think
//also will need to add a custom CORS config kida like e.Use(middeware.CustomCORS(http://exampleurl.com))
//prob for dev e.Use(middeware.CORS()) will just allow any origin for any request
func (g *GoLive)Use(port string){  //i think pointer will be needed as i want global middeware
	//will need to take the contect function somehow to access the this will need more things like maybe pass a specific port for middeware
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*") pass a port here maybe
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
//wanna keep this comment just to know what the custom method represents
func (g *GoLive) GET(path string, /*mux *http.ServeMux,*/ handle FunctionHandler) { //get request wrapper for simple usage
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return
	}
	fullGetPath := fmt.Sprintf("GET %s", path)
  g.Mux.HandleFunc(fullGetPath, func(w http.ResponseWriter, r *http.Request) {
    // if r.Method != http.MethodGet {
    //   http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    //   return
    // }

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
func (g *GoLive)POST(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){ //put request wrapper
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return
	}
	fullPostPath:= fmt.Sprintf("POST %s", path)
  g.Mux.HandleFunc(fullPostPath, func(w http.ResponseWriter, r *http.Request) {
    // if r.Method != http.MethodPost{
    //   http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    //   return
    // }

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
func (g *GoLive)DELETE(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){ //DELETE request wrapper
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return// may need to add this to the others
	}
	fullDeletePath := fmt.Sprintf("DELETE %s", path)
  g.Mux.HandleFunc(fullDeletePath, func(w http.ResponseWriter, r *http.Request) {
    // if r.Method != http.MethodDelete{
    //   http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    //   return
    // }

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
func (g *GoLive)PUT(path string, /*mux *http.ServeMux,*/ handle FunctionHandler){ //PUT request wrapper
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return// may need to add this to the others
	}
	fullPutPath := fmt.Sprintf("PUT %s", path)
  g.Mux.HandleFunc(fullPutPath, func(w http.ResponseWriter, r *http.Request) {
    // if r.Method != http.MethodPut{
    //   http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    //   return
    // }

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

// type GoLive struct{
// 	Mux *http.ServeMux
// }

//this function is what starts the server should be put at the end of the main file
func (g *GoLive)StartServer(port string){
	server := &http.Server{
		Addr:    port,
		Handler: middleware.Logging(g.Mux), //this is where the output for Requests are
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
	fmt.Println("--------------------------------------------------")

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
	} else if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
	}
}
