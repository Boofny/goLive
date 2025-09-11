package main

import (
	"GoLive/routes"
	"net/http"
)

func main() {
	e := routes.New()
	
	mux := http.NewServeMux()
	e.GET("GET /tester", mux, func (w http.ResponseWriter, r *http.Request){
		resp2 := map[string]any{
			"david": "New function",
		}
		routes.SendJSON(w, http.StatusAccepted, resp2)
	})


	mux.HandleFunc("POST /pong", e.PostJsonDefault)

	routes.StartServer(":8080", mux)
}


func GetStringDefault(w http.ResponseWriter, r *http.Request){
	resp := map[string]any{
		"name": "Hello world",
	}
	routes.SendJSON(w, http.StatusAccepted, resp)
}
