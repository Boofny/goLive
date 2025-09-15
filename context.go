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

func (c *Context) ERROR(/*w http.ResponseWriter,*/ status int, errorMsg string) error {
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte(errorMsg))
	if err != nil { //idk 
		return err
	}
	return nil
}


