package goLive

import (
	"encoding/json"
	"net/http"
)

type Context struct{
	Writer http.ResponseWriter
	Request *http.Request
}

func (c *Context) JSON(/*w http.ResponseWriter,*/ status int, data any) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	return json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) STRING(/*w http.ResponseWriter,*/ status int, data string)error{
  c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte(data))
	return err 
}
//check these fucntions aswell

func (c *Context) ERROR(/*w http.ResponseWriter,*/ status int, errorMsg string) error {
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte(errorMsg))
	return err
}

// func (c *Context) REDIRECT(status int, redirectUrl string)error{
// 	//feel this could be done better
// 	http.Redirect(c.Writer, c.Request, redirectUrl, status)
// 	// c.Writer.WriteHeader(200)
// 	return nil
// }
//
func (c *Context) REDIRECT(status int, redirectUrl string) error {
	if status < 300 || status > 308 {
		return ErrInvalidRedirectCode
	}
	http.Redirect(c.Writer, c.Request, redirectUrl, status)
	c.Writer.WriteHeader(status)
	return nil
}
// func (c *Context) Error(err error) {
// 	c.echo.HTTPErrorHandler(err, c)
// }

// func (c *context) Echo() *Echo {
// 	return c.echo
// }

