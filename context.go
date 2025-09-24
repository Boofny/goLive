/*
GoLive http frame work this is the Context file that holds 
varias function to send read and detect errors
most of these ,methods will need a http status code 200 404 etc...
*/

package goLive

import (
	"encoding/json"
	"io"
	"net/http"
)

//custom Context struct to have http types tied to methods rather than passed to functions
type Context struct{
	Writer http.ResponseWriter
	Request *http.Request
}

//when a request needs json to be send this function is used taking a http status code and any for of data mainly maps
func (c *Context) SendJson(status int, data any) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	return json.NewEncoder(c.Writer).Encode(data)
}

//used for when the http method sends a json to the server and need to extract json datad
//dev must pass the data by address when using in order to bind to the models that are defined
func (c *Context)ReadJson(data any)error{
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	defer c.Request.Body.Close()
	return json.Unmarshal(body, data)
}

func (c *Context)Valid(status int)error{
	if status < 200 || status > 399{
		return ErrInvalidRedirectCode
	}
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte("Valid"))
	return err
}

//sends a simple text only string to the client good for fast tests 
func (c *Context) SendString(status int, data string)error{
  c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte(data))
	return err 
}

//sends error into the request for more custom and simple error handling in the http methods
func (c *Context) Error(status int, errorMsg string) error {
	if status < 400 || status > 599{
		return ErrInvalidRedirectCode
	}
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Writer.WriteHeader(status)
	_, err := c.Writer.Write([]byte(errorMsg))
	return err
}

//redirects the clients request to the needed url or other path
func (c *Context) Redirect(status int, redirectUrl string) error {
	if status < 300 || status > 308 {
		return ErrInvalidRedirectCode
	}
	http.Redirect(c.Writer, c.Request, redirectUrl, status)
	c.Writer.WriteHeader(status)
	return nil
}

func (c *Context)Param(data string)string{ //for now this works with only string
	foundData := c.Request.PathValue(data)
	return foundData
}

func (c *Context)QueryGet(data string)string{
	foundQuery := c.Request.URL.Query().Get(data)
	return foundQuery
}




