package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GlobalMiddleware(h http.Handler) http.Handler {
	// Middleware and mux.MiddlewareFunc types are different.
	// This is wy we should wrap h handler and convert it to http.HandlerFunc
	var wrapped http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) { h.ServeHTTP(w, r) }

	// this middlewares will be called on every request to every handler
	return ApplyMiddleware(wrapped, RequestTime, UrlPath)
}

func main() {
	r := mux.NewRouter()

	r.Use(GlobalMiddleware)

	r.HandleFunc("/", root)
	r.HandleFunc("/login", ApplyMiddleware(signIn, Authenticate)).Methods("POST")
	r.HandleFunc("/register", register).Methods("POST")
	// one way to apply middleware
	r.HandleFunc("/users", Authorize(users))
	// another way to apply same middleware
	r.HandleFunc("/subscriptions", ApplyMiddleware(subscriptions, Authorize))
	http.ListenAndServe("localhost:8080", r)
}
