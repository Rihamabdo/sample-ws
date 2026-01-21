package middleware

import (
	"context"
	"net/http"
	"strings"

	"sample-ws/internal/utils"
)

type ctxkey string

const UsernameKey ctxkey = "username"

func JWTAuth(secret []byte, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" || !strings.HasPrefix(auth, "Bearer") {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		t, claims, err := utils.ParseToken(secret, tokenStr)
		if err != nil || !t.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		username, ok := claims["username"].(string)
		if !ok || username == "" {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UsernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
