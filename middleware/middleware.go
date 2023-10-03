package middleware

import "net/http"

// This type describes handler that can decorate another handler:
// decoratedHandler := Middleware(handler)
type Middleware func(fn http.HandlerFunc) http.HandlerFunc

// This function takes handler and a list of decorators and apply them
// one by one: Apply(handler, one, two) => two(one(handler))
func Apply(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, next := range middlewares {
		handler = next(handler)
	}
	return handler
}
