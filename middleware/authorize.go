package middleware

import (
	"errors"
	"log"
	"net/http"
)

func Authorize() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {

			cookie, err := req.Cookie(SESSION_COOKIE)

			if err != nil {
				switch {
				case errors.Is(err, http.ErrNoCookie):
					http.Error(w, "You are not logged in", http.StatusBadRequest)
				default:
					log.Println(err)
					http.Error(w, "server error", http.StatusInternalServerError)
				}
				return
			}

			if !authenticated[cookie.Value] {
				http.Error(w, "Your session is expired", http.StatusUnauthorized)
				return
			}

			next(w, req)
		}
	}
}
