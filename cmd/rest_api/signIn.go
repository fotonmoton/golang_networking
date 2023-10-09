package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func signIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	var newAttempt attempt

	err = json.Unmarshal(body, &newAttempt)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	originPassword, loginExists := registered[newAttempt.Login]
	if !loginExists {
		http.Error(w, "login doesn't exists", http.StatusBadRequest)
		return
	}

	h := sha1.New()
	h.Write([]byte(newAttempt.Password))
	guessPassword := password(hex.EncodeToString(h.Sum(nil)))

	if originPassword != guessPassword {
		http.Error(w, "Password doesn't match", http.StatusBadRequest)
		return
	}

	// we modify request with new context value
	// to notify that this request is authenticated
	*r = *r.Clone(context.WithValue(r.Context(), CONTEXT_AUTH_KEY, true))
}
