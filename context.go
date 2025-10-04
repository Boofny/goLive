/*
GoLive http frame work this is the Context file that holds
varias function to send read and detect errors
most of these ,methods will need a http status code 200 404 etc...
*/

// TODO: make error and valid send json make the key for the map the param

package goLive

import (
	"encoding/json"
	"errors"

	// "io"
	"net/http"
)

//custom Context struct to have http types tied to methods rather than passed to functions
type Context struct{
	Writer http.ResponseWriter
	Request *http.Request
}

//when a request needs json to be send this function is used taking a http status code and any form of data mainly maps
func (c *Context) SendJSON(status int, data any) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	return json.NewEncoder(c.Writer).Encode(data)
}

func (c *Context) ReadJSON(data any) error {
	defer c.Request.Body.Close()

	if c.Request.Header.Get("Content-Type") != "application/json" {
		return errors.New("invalid content type")
	}

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // optional: prevents extra fields in JSON
	return decoder.Decode(data)
}

// func (c *Context)ReadJSON(data any)error{
// 	body, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		return err
// 	}
// 	defer c.Request.Body.Close()
// 	return json.Unmarshal(body, data)
// }

//sends a simple text only string to the client good for fast tests json key is "response 
func (c *Context) SendSTRING(status int, data string)error{
	resp := map[string]string{
		"response": data,
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	err := json.NewEncoder(c.Writer).Encode(resp)
	return err
}

//Sends a json response with custom message json key is "response"
func (c *Context)Valid(status int, validMsg string)error{
	resp := map[string]string{
		"response": validMsg,
	}
	if status < 200 || status > 399{
		return ErrInvalidRedirectCode
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	err := json.NewEncoder(c.Writer).Encode(resp)
	return err
}

//sends error into the request for more custom and simple error handling in the http methods
func (c *Context) Error(status int, errorMsg string) error {
	// _, err := c.Writer.Write([]byte(errorMsg))
	resp := map[string]string{
		"response": errorMsg,
	}
	if status < 400 || status > 599{
		return ErrInvalidRedirectCode
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)
	err := json.NewEncoder(c.Writer).Encode(resp)
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

//gets the value of path url param
func (c *Context)Param(data string)string{
	foundData := c.Request.PathValue(data)
	return foundData
}

//gets the Query from the url ?=
func (c *Context)QueryGet(data string)string{
	foundQuery := c.Request.URL.Query().Get(data)
	return foundQuery
}

// TODO: need to decide if these functions are part of Context or GoLive struct and how they fit in the framework
func (c *Context)ReciveFile(){

}

func (c *Context)SendFile(filepath string)error{
	http.ServeFile(c.Writer, c.Request, filepath)
	return nil
}


