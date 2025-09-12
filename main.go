package main

import (
	"GoLive/routes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	e := routes.New()

	server := e.MuxHandle() //i wanna see if i really need this so i dont have to pass server to the GET and POST functions
	e.GET("/pers", &server, func(w http.ResponseWriter, r *http.Request) error{
		resp := map[string]any{
			"nerd": "nananana",
		}
		return e.SendJSON(w, http.StatusOK, resp)
	})

	e.POST("/ping", &server, func(w http.ResponseWriter, r *http.Request)error{
		type User struct{
			Name string `json:"name"`
			Age int `json:"age"`
		}
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Error in post", http.StatusBadRequest)
			return err
		}
		ret := fmt.Sprintf("Hello %s", user.Name)
		code := map[string]any {"resp": ret, "resp2": user.Age}
		return e.SendJSON(w, http.StatusOK, code)
	})

	e.StartServer(":8080", &server)
}

