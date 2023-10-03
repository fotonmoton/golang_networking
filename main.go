package main

import (
	"net/http"

	"github.com/fotonmoton/golang_networking/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(func(h http.Handler) http.Handler {
		// middleware.Middleware and mux.MiddlewareFunc types are different.
		// This is wy we should wrap h handler and convert it to http.HandlerFunc
		var wrapped http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) { h.ServeHTTP(w, r) }

		// this middlewares will be called on every request
		return middleware.Apply(wrapped, middleware.RequestTime(), middleware.UrlPath())
	})

	r.HandleFunc("/", root)
	// one way to apply middleware
	r.HandleFunc("/login", middleware.Authenticate()(signIn)).Methods("POST")
	r.HandleFunc("/register", register).Methods("POST")
	// another way to apply middleware
	r.HandleFunc("/users", middleware.Apply(users, middleware.Authorize()))
	r.HandleFunc("/subscriptions", middleware.Apply(subscriptions, middleware.Authorize()))
	http.ListenAndServe("localhost:8080", r)
}
