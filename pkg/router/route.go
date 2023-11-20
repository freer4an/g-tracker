package router

import "net/http"

type Route struct {
	Method      string
	Pattern     string
	Query       string
	middlewares []MiddlewareFunc
	Handler     http.HandlerFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func (r *Route) Use(middlewares ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, middlewares...)
}
