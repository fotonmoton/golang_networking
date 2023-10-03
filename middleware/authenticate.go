package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func Authenticate() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			next(w, req)

			// Wrapped handler should modify context
			// and set CONTEXT_AUTH_KEY key to true.
			// Otherwise we do not authenticate this request.
			// Also, next(w, req) should not write to the body
			// because SetCookie will write to header and it will be
			// useless
			isAuthenticated, exists := req.Context().Value(CONTEXT_AUTH_KEY).(bool)

			if !exists || !isAuthenticated {
				return
			}

			session := uuid.NewString()

			authenticated[session] = true

			cookie := http.Cookie{
				Name:     SESSION_COOKIE,
				Value:    session,
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteNoneMode,
			}

			http.SetCookie(w, &cookie)

			w.Write([]byte("you are logged in"))
		}
	}
}
