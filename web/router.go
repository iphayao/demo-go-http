package web

import (
	"net/http"
)

type Router interface {
	Routes
}

type Routes interface {
	GET(string, HandlerFunc)
}

type Route struct {
	Path    string
	Method  string
	Handler HandlerFunc
	engine  *Engine
	root    bool
}

func (r *Route) GET(path string, h HandlerFunc) {
	r.handler(http.MethodGet, path, h)
}

func (r *Route) handler(method string, path string, h HandlerFunc) {
	r.engine.addRoute(method, path, h)
}
