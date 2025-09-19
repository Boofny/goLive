package goLive

import (
	"encoding/json"
	"net/http"
)

type Context struct{
	Writer http.ResponseWriter
	Request *http.Request
}

func (c *Context) JSON(status int, data any) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	return json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) STRING(status int, data string)error{
  c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte(data))
	return err 
}

func (c *Context) ERROR(status int, errorMsg string) error {
	if status < 400 || status > 599{
		return ErrInvalidRedirectCode
	}
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte(errorMsg))
	return err
}

func (c *Context) REDIRECT(status int, redirectUrl string) error {
	if status < 300 || status > 308 {
		return ErrInvalidRedirectCode
	}
	http.Redirect(c.Writer, c.Request, redirectUrl, status)
	c.Writer.WriteHeader(status)
	return nil
}

//need to make a bind function aka readJson from post req 
		// err := c.Bind(&link)
		// if err != nil{
		// 	return c.JSON(http.StatusBadRequest, map[string]string{
		// 		"error": "Invalid json",
		// 	})
		// }
