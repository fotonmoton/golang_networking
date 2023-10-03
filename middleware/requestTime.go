package middleware

import (
	"log"
	"net/http"
	"time"
)

// This middleware logs request's time
func RequestTime() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			next(w, req)
			log.Println(req.URL.Path, "takes:", time.Since(start))
		}
	}
}
