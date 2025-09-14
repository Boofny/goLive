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
	c := goLive.Tools()

	mid := http.NewServeMux()

	e.GET("/hello", mid, func(w http.ResponseWriter, r *http.Request) error {
		return c.STRING(w, 200, "Hello world")
	})

	e.StartServer(":8080", mid)
}
```

##Clone
Get started by cloning this repo
```bash
git clone https://github.com/Boofny/golive.git
```
