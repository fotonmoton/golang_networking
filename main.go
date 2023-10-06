package main

import (
	"net/http"

	"github.com/fotonmoton/golang_networking/files"
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
	f.HandleFunc("", files.ListFiles).Methods("GET")
	f.HandleFunc("", files.CreateFile).Methods("POST")
	f.HandleFunc("/{file}", files.ShowFile).Methods("GET")
	f.HandleFunc("/{file}", files.UpdateFile).Methods("PUT")
	f.HandleFunc("/{file}", files.DeleteFile).Methods("DELETE")
	http.ListenAndServe("localhost:8080", r)
}
