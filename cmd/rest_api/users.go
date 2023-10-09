package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var usersDB = []map[string]any{{"name": "Greg"}, {"name": "John"}}

var users http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	log.Printf("users handler")

	w.WriteHeader(http.StatusOK)
	response, _ := json.MarshalIndent(usersDB, "", " ")
	w.Write(response)
}
