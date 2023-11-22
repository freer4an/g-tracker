package router

import (
	"fmt"
	"net/http"
	"strings"
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

// type HttpLogger func(http.HandlerFunc) http.HandlerFunc
func NewRouter() *Router {
	return &Router{
		routes: []*Route{},
	}
}

// Add adds a new route with the specified method, pattern and handler.
func (r *Router) AddRoute(method, pattern string, handler http.HandlerFunc) *Route {
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

func (r *Router) ServeStatic(pattern string, handler http.Handler) {
	r.AddRoute(MethodGet, pattern, handler.ServeHTTP).isStatic = true
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.isStatic && strings.HasPrefix(req.URL.Path, route.Pattern) {
			route.Handler.ServeHTTP(w, req)
			return
		}
		// re := regexp.MustCompile(route.Pattern)
		if route.Method == req.Method && (route.Pattern == req.URL.Path) {
			// for _, middleware := range route.middlewares {
			// 	route.Handler = middleware(route.Handler)
			// }
			route.Handler.ServeHTTP(w, req)
			return
		}
	}
	response := fmt.Sprintf("%v - %s", http.StatusNotFound, http.StatusText(http.StatusNotFound))
	http.Error(w, response, http.StatusNotFound)
}

// func (r *Router) AddLogger(handler HttpLogger) {
// 	for _, route := range r.routes {
// 		if route.middlewares != nil {
// 			panic("Logger should be used first and once only")
// 		}
// 		route.Use(MiddlewareFunc(handler))
// 	}
// }
