package main

import (
	"net/http"
	"github.com/Boofny/goLive"
)

func main() {
	e := goLive.New()

	e.GET("/hello", func(c *goLive.Context) error {
		id := c.QueryGet("name")
		if id != "David"{
			return c.Error(http.StatusNotFound, "Error here twin")
		}
		resp := map[string]string{
			"Name": id,
		}
		return c.SendJson(http.StatusOK, resp)
	})

	e.GET("/redi/{name}", func(c *goLive.Context) error {
		url := "https://github.com/Boofny"
		// name := c.Request.PathValue("name")
		name := c.Param("name")

		if name != "devCode"{
			return c.Error(http.StatusNotFound, "Error in /redi")
		}else{
			return c.Redirect(http.StatusFound, url) 
		}
	})

	e.POST("/read", func(c *goLive.Context) error {

		type User struct{ //should be from models dir
			Name string `json:"name"`
			Id int `json:"id"`
			Email string `json:"email"`
		}

		var data User
		err := c.ReadJson(&data)
		if err != nil {
			return c.Error(http.StatusNotFound, "Error in /read")
		}

		return c.SendJson(http.StatusOK, map[string]any{
			"nameresp": "Hello " + data.Name,
			"idresp": data.Id,
			"emailresp": data.Email,
		})

	})

	e.StartServer(":8080")
}
