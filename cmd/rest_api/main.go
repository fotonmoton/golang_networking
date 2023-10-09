package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

	// Files REST API
	f := r.PathPrefix("/files").Subrouter()
	f.HandleFunc("", ListFiles).Methods("GET")
	f.HandleFunc("", CreateFile).Methods("POST")
	f.HandleFunc("/{file:.+}", ShowFile).Methods("GET")
	f.HandleFunc("/{file:.+}", UpdateFile).Methods("PUT")
	f.HandleFunc("/{file:.+}", DeleteFile).Methods("DELETE")
	log.Println("listening on localhost:8080")
	http.ListenAndServe("localhost:8080", r)
}
