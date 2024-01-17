package router

import (
	"fmt"
	"net/http"
	"strings"
)

type routeParam struct {
	key   string
	value string
	index int
}

type RouteParams []routeParam

type Route struct {
	Method     string
	Pattern    string
	patternSlc []string
	RouteParams
	middlewares []MiddlewareFunc
	Handler     http.Handler
	isStatic    bool
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func (r *Route) Use(middlewares ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func newRoute(method, pattern string,
	patternSlc []string,
	routeParams []routeParam,
	handler http.HandlerFunc) *Route {
	if method == "" {
		panic("invalid method")
	}
	if pattern == "" {
		panic("invalid pattern")
	}
	return &Route{
		Method:      method,
		Pattern:     pattern,
		patternSlc:  patternSlc,
		RouteParams: routeParams,
		Handler:     handler,
	}
}

func (r *Route) matchRoute(method, path string) bool {
	// absolute case
	if r.Method == method && r.Pattern == path {
		return true
	}

	// relative case
	if len(r.RouteParams) == 0 {
		return false
	}
	patternSlc := strings.Split(path, "/")
	if len(r.patternSlc) != len(patternSlc) {
		return false
	}
	for i := 0; i < len(patternSlc); i++ {
		if patternSlc[i] == r.patternSlc[i] {
			continue
		}
		if r.patternSlc[i] != ":" {
			return false
		}
		for j := 0; j < len(r.RouteParams); j++ {
			if r.RouteParams[j].index == i {
				r.RouteParams[j].value = patternSlc[i]
				break
			}
		}
	}
	return true
}

func (r *Route) Describe() {
	fmt.Printf("%s \t %s \tParams: %v\n", r.Method, r.Pattern, r.RouteParams)
}

func (params *RouteParams) GetParam(key string) string {
	for _, param := range *params {
		if param.key == key {
			return param.value
		}
	}
	return ""
}
