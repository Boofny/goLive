package main

import (
	"net/http"

	"github.com/Boofny/goLive"
	"github.com/Boofny/goLive/middleware"
)

func main() {
	e := goLive.Launch()

	e.Chain(
		middleware.CustomCORS(AllowOrigins []string{"*"}),
		middleware.Logger(),
	)

	// example of a get request with ?=
	e.GET("/hello", func(c *goLive.Context) error {
		id := c.QueryGet("name")
		if id != "David"{
			return c.Error(http.StatusNotFound, "Error here twin")
		}
		resp := map[string]string{
			"Name": id,
		}
		return c.SendJSON(http.StatusOK, resp)
	})

	//example of a get request with path values
	e.GET("/redi/{name}", func(c *goLive.Context) error {
		url := "https://github.com/Boofny"
		name := c.Param("name")

		if name != "devCode"{
			return c.Error(http.StatusNotFound, "Error in /redi")
		}else{
			return c.Redirect(http.StatusFound, url) 
		}
	})
	
	//example of reading json from post request
	e.POST("/read", func(c *goLive.Context) error {

		type User struct{ 
			Name string `json:"name"`
			Id int `json:"id"`
			Email string `json:"email"`
		}

		var data User
		err := c.ReadJSON(&data)
		if err != nil {
			return c.Error(http.StatusNotFound, "Error in /read")
		}

		return c.SendJSON(http.StatusOK, map[string]any{
			"nameresp": "Hello " + data.Name,
			"idresp": data.Id,
			"emailresp": data.Email,
		})

	})

	v1 := e.GroupRoutes("/v1")

	v1.GET("/ping", func(c *goLive.Context) error {
		return c.SendSTRING(http.StatusOK, "pong v1")
	})

	e.StartServer(":8080")
}
