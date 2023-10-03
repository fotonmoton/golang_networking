package middleware

import (
	"log"
	"net/http"
)

// This middleware logs paths of incoming requests
func UrlPath() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			log.Println("current request path is", req.URL.Path)
			next(w, req)
		}
	}
}
