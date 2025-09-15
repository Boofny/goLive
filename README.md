## GoLive!
### Fast web frame work for golang built on std lib http

```go
package main

import (
	"net/http"
	"github.com/Boofny/goLive"
)

func main() {
	e := goLive.New()

	e.GET("/hello", func(c *goLive.Context)error{
		return c.STRING(http.StatusOK, "Hello world") //send out your data
	})

	e.StartServer("8080")
}
```

## Installation 
### Get started by downloading this module
```bash
go get github.com/Boofny/goLive@latest
```
