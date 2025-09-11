package main

import (
	"GoLive/routes"
	"encoding/json"
	"net/http"
)

func main() {
	e := routes.New()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", e.GetJsonDefault)
	mux.HandleFunc("POST /pong", e.PostJsonDefault)
	mux.HandleFunc("GET /nerd", func(w http.ResponseWriter, r *http.Request){
		resp := map[string]any{
			"nerd": "nananana",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(resp)
	})

	routes.StartServer(":8080", mux)
}

