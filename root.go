package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("Served request")
	log.Println(r.Header.Get("User-Agent"))
	w.WriteHeader(http.StatusBadRequest)
	response, _ := json.Marshal(map[string]any{"hello": "123"})
	w.Write(response)
}
