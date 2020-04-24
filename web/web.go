package web

import (
	"log"
	"net/http"
	"sync"
)

type HandlerFunc func(Context) error

type Engine struct {
	Route
	pool     sync.Pool
	routes   map[string]Route
}

func New() *Engine {
	e := &Engine{
		Route: Route{
			Handler: nil,
			Path:    "/",
			root:    true,
		},
	}

	e.Route.engine = e
	e.routes = make(map[string]Route)
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}

	return e
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := e.pool.Get().(*Context)
	c.request = r
	c.response = w

	e.handleHTTPRequest(*c)
	e.pool.Put(c)
}

func (e *Engine) addRoute(method string, path string, handle HandlerFunc) {
	route := Route{
		Path:    path,
		Method:  method,
		Handler: handle,
		engine:  e,
		root:    false,
	}

	e.routes[method+path] = route
}

func (e *Engine) Start(address string) error {
	log.Printf("HTTP server started on %s", address)
	return http.ListenAndServe(address, e)
}

func (e *Engine) handleHTTPRequest(c Context) {
	method := c.request.Method
	path := c.request.URL.Path

	route := e.routes[method+path]
	if err := route.Handler(c); err != nil {
		log.Printf("Error %s", err)
	}
}

func (e *Engine) allocateContext() *Context {
	return &Context{engine: e, KeysMutex: &sync.RWMutex{}}
}


