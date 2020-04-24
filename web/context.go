package web

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Context struct {
	request  *http.Request
	response http.ResponseWriter

	engine    *Engine
	KeysMutex *sync.RWMutex
}

func (c *Context) JSON(code int, value interface{}) error {
	return c.Render(code, value)
}

func (c *Context) String(status int, message string) error {
	return c.JSON(status, message)
}

func (c *Context) Render(code int, value interface{}) error {
	enCode := json.NewEncoder(c.response)

	c.Status(code)
	return enCode.Encode(value)
}

func (c *Context) Status(code int) {
	c.response.WriteHeader(code)
}
