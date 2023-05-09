package sin

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type H map[string]interface{}


// sin.Context, http.ResponseWriter + http.Request
type Context struct {
	Writer  http.ResponseWriter
	Req *http.Request
	
	Path string //  request path
	Method string  // response method
	Params map[string]string

	StatusCode int // response statuscode
}

func(c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// create new context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context {
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}

// Get PostForm info by key
func (c *Context) PostForm(key string) string  {
	return c.Req.FormValue(key)
}

// Get url info by key
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Response Statuscode
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// Respones Header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key,value)
}

// response -- string
// Set response head -> Set StatusCode -> Set response body
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...))) // write data into response body
}

// response -- JSON
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err  != nil {
		http.Error(c.Writer, err.Error(),500)
	}
}

// response -- data
func (c *Context)  Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// response -- HTML
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
