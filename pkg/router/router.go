package router

import (
	"fmt"
	"net/http"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

type Router struct {
	routes []*Route
}

type HttpLogger func(http.HandlerFunc) http.HandlerFunc

func NewRouter() *Router {
	return &Router{
		routes: []*Route{},
	}
}

// Add adds a new route with the specified method, pattern and handler.
func (r *Router) Add(method, pattern string, handler http.HandlerFunc) *Route {
	if method != MethodGet && method != MethodPost && method != MethodPut && method != MethodDelete {
		panic("invalid method")
	}
	route := &Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	}
	r.routes = append(r.routes, route)
	return route
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && route.Pattern == req.URL.Path {
			for _, middleware := range route.middlewares {
				route.Handler = middleware(route.Handler)
			}
			route.Handler.ServeHTTP(w, req)
			return
		}
	}
	response := fmt.Sprintf("%v - %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	http.Error(w, response, http.StatusNotFound)
}

func (r *Router) AddLogger(handler HttpLogger) {
	for _, route := range r.routes {
		if route.middlewares != nil {
			panic("Logger should be used first")
		}
		route.Use(MiddlewareFunc(handler))
	}
}

func (r *Router) Routes() {
	for _, route := range r.routes {
		fmt.Printf("%+v\n", route)
	}
}
