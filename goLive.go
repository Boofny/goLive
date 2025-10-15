package goLive

import (
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

//defined struct for starting server and chaining middleware
type GoLive struct{
	Mux *http.ServeMux
	middlewares []middleware.Middleware
}

//Method for starting the goLive session
func Launch()*GoLive{ 
	return &GoLive{
		Mux: http.NewServeMux(),
	}
}

//wanna keep this comment just to know what the custom method represents
func (g *GoLive) GET(path string, /*mux *http.ServeMux,*/ handle FunctionHandler) { //get request wrapper for simple usage
	if path == "/favicon.ico" { //just ignore this will prob redirect in future
  	return
	}
	fullGetPath := fmt.Sprintf("GET %s", path)
  g.Mux.HandleFunc(fullGetPath, func(w http.ResponseWriter, r *http.Request) {

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

//To serve a static file html txt png etc
func (g *GoLive)ServeStatic(urlPath, filepath string)error{
	_, err := os.Stat(filepath)
	if os.IsNotExist(err){
		return fmt.Errorf("File does not exist %s", filepath)
	}

  g.Mux.HandleFunc(urlPath, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
		}
		http.ServeFile(w, r, filepath)
  })
	return nil
}

//To serve an entire dir
func (g *GoLive)ServeDir(urlPath, dirPath string)error{
	fs := http.FileServer(http.Dir(dirPath))
	g.Mux.Handle(urlPath, http.StripPrefix(urlPath, fs))

	return nil
}

//use passes a variadic value of Middleware that is appended to the g.middlewares slice
func (g *GoLive)Chain(mw ...middleware.Middleware){
	g.middlewares = append(g.middlewares, mw...)
}

func (g *GoLive)GroupRoutes(urlPath string, m ...middleware.Middleware){
	//this will need route groiping and middlewares chainging
}

//this function is what starts the server should be put at the end of the main file
func (g *GoLive)StartServer(port string){

	stack := middleware.CreateStack(g.middlewares...)

	server := &http.Server{
		Addr:    port,
		Handler: stack(g.Mux), //where g.Mux is added after middleware chaining 
		// Handler: middleware.Logging(g.Mux), //this is where the output for Requests are
	}

	startingDisaply(port)

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Server closed")
	} else if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
	}
}

func startingDisaply(port string){
	banner :=  `
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
	fmt.Println(blue, banner)
	fmt.Print("\033[34m >>> \033[0m")
	fmt.Print("Server started successfully on port " +  yellow + port + reset)
	fmt.Println("\033[34m <<< \033[0m")
	fmt.Println("--------------------------------------------------")
}
