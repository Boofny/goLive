package main

import (
	"GoLive/routes"
	"net/http"
)

func main() {
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", routes.GetJsonDefault)
	// mux.HandleFunc("GET /ping", routes.GetStringDefault)
	mux.HandleFunc("POST /pong", routes.PostJsonDefault)

	routes.StartServer(":8080", mux)
}
