package main

import (
	"GoLive/routes"
	"net/http"
)

func main() {
	e := routes.New()
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", e.GetJsonDefault)
	mux.HandleFunc("POST /pong", e.PostJsonDefault)

	routes.StartServer(":8080", mux)
}
