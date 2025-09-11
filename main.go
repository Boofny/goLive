package main

import (
	"GoLive/routes"
	"net/http"
)

func main() {
	e := routes.New()
	e.GetStringDefault()

	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", routes.GetJsonDefault)
	mux.HandleFunc("POST /pong", routes.PostJsonDefault)

	routes.StartServer(":8080", mux)
}
