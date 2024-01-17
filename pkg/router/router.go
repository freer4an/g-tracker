package router

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

const (
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
)

var paramsKey = struct{}{}

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
func (r *Router) addRoute(method, pattern string, handler http.HandlerFunc) *Route {
	// if method != MethodGet && method != MethodPost && method != MethodPut && method != MethodDelete {
	// 	panic("invalid method")
	// }
	var routeParams []routeParam

	patternSlc := strings.Split(pattern, "/")
	for i := 0; i < len(patternSlc); i++ {
		var rParam routeParam
		if len(patternSlc[i]) == 0 {
			continue
		}
		if patternSlc[i][0] == ':' {
			rParam.key = patternSlc[i][1:]
			if len(rParam.key) == 0 {
				panic("invalid pattern")
			}
			rParam.index = i
			routeParams = append(routeParams, rParam)
			patternSlc[i] = ":"
		}
	}
	pattern = strings.Join(patternSlc, "/")
	route := newRoute(method, pattern, patternSlc, routeParams, handler)
	r.routes = append(r.routes, route)
	route.Describe()
	return route
}

func (r *Router) Get(pattern string, handler http.HandlerFunc) *Route {
	return r.addRoute(MethodGet, pattern, handler)
}

func (r *Router) Post(pattern string, handler http.HandlerFunc) *Route {
	return r.addRoute(MethodPost, pattern, handler)
}

func (r *Router) Put(pattern string, handler http.HandlerFunc) *Route {
	return r.addRoute(MethodPut, pattern, handler)
}

func (r *Router) Delete(pattern string, handler http.HandlerFunc) *Route {
	return r.addRoute(MethodDelete, pattern, handler)
}

func (r *Router) ServeStatic(pattern string, handler http.Handler) {
	r.addRoute(MethodGet, pattern, handler.ServeHTTP).isStatic = true
}

// ServeHTTP implements the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.isStatic && strings.HasPrefix(req.URL.Path, route.Pattern) {
			route.Handler.ServeHTTP(w, req)
			return
		}
		// TODO: QUERY detector RouteMatcher
		// re := regexp.MustCompile(route.Pattern)
		// if (route.Method == req.Method && (route.Pattern == req.URL.Path)) || (route.RouteQuery != "" && re.MatchString(req.URL.Path)) {
		// 	// for _, middleware := range route.middlewares {
		// 	// 	route.Handler = middleware(route.Handler)
		// 	// }
		// 	if route.RouteQuery != "" {
		// 		i := strings.Index(req.URL.Path, route.RouteQuery)
		// 		fmt.Println(req.URL.Path, i)
		// 		return
		// 		ctx := context.WithValue(req.Context(), "id", req.URL.Path[i:])
		// 		req = req.WithContext(ctx)
		// 		fmt.Println(req.URL.Path[i:])
		// 	}
		// 	route.Handler.ServeHTTP(w, req)
		// 	return
		// }
		if !route.matchRoute(req.Method, req.URL.Path) {
			continue
		}
		if len(route.RouteParams) == 0 {
			route.Handler.ServeHTTP(w, req)
			return
		}
		ctx := context.WithValue(req.Context(), paramsKey, route.RouteParams)
		route.Handler.ServeHTTP(w, req.WithContext(ctx))
		return
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

// func (r *Router) AddLogger(handler HttpLogger) {
// 	for _, route := range r.routes {
// 		if route.middlewares != nil {
// 			panic("Logger should be used first and once only")
// 		}
// 		route.Use(MiddlewareFunc(handler))
// 	}
// }

func paramsList(ctx context.Context) RouteParams {
	p, _ := ctx.Value(paramsKey).(RouteParams)
	return p
}

func GetParam(ctx context.Context, key string) string {
	p := paramsList(ctx)
	for _, param := range p {
		if param.key == key {
			return param.value
		}
	}
	return ""
}

func GetParamInt(ctx context.Context, key string) int {
	p := paramsList(ctx)
	for _, param := range p {
		if param.key == key {
			i, _ := strconv.Atoi(param.value)
			return i
		}
	}
	return 0
}
