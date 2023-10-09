package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var subscriptionsDB = []map[string]any{{"login1": []string{"sub1", "sub1"}}}

func subscriptions(w http.ResponseWriter, r *http.Request) {
	log.Printf("users handler")

	w.WriteHeader(http.StatusOK)
	response, _ := json.MarshalIndent(subscriptionsDB, "", " ")
	w.Write(response)
}
