package middleware

import (
	"net/http"
	"crypto/sha256"
	"crypto/subtle"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedU := sha256.Sum256([]byte("admin"))
			expectedP := sha256.Sum256([]byte("admin"))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:],expectedU[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:],expectedP[:]) == 1)

			if usernameMatch == passwordMatch {
				next.ServeHTTP(w,r)
				return
			}

		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
