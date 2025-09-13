package main

import (
	"GoLive/routes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	e := routes.New()

	server := http.NewServeMux()

	e.GET("/romdd", server, func(w http.ResponseWriter, r *http.Request) error {
		resp := map[string]any{
			"name": "nanananana",
		}
		return e.SendJSON(w, http.StatusOK, resp)
	})

	e.GET("/test", server, func(w http.ResponseWriter, r *http.Request)  error{
		query := r.URL.Query()
		resp := map[string]any{
			"id": query.Get("id"),
			"p": query.Get("p"),
			"page": query.Get("page"),
		}
		idConv, err := strconv.Atoi(resp["id"].(string))
		if err != nil {
			return err
		}
		if idConv != 1323{
			return e.SendJSON(w, http.StatusBadRequest, "Error in /test")
		}
		return e.SendJSON(w, http.StatusOK, resp)
	})

	e.POST("/ping", server, func(w http.ResponseWriter, r *http.Request)error{
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

	e.StartServer(":8080", server)
}

