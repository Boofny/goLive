## GoLive!
### Fast web frame work for golang built on std lib http

## Installation 
### Get started by downloading this module
#### Need Golang 1.25.1+
```bash
go get github.com/Boofny/goLive@latest
```

```go
package main

import (
	"net/http"
	"github.com/Boofny/goLive"
	"github.com/Boofny/goLive/middleware"
)

func main() {
	e := goLive.Launch()

	e.Chain(
        middleware.CORS(),
        middleware.Logger(),
	)

	e.GET("/hello", func(c *goLive.Context)error{
		return c.SendSTRING(http.StatusOK, "Hello World") //send out your data
	})

	e.StartServer(":8080")
}
```

## Try it out!
### In the terminal try.
```bash
curl -X GET localhost:8080/hello
```
### Or in your browser 
```bash
http://localhost:8080/hello
```
