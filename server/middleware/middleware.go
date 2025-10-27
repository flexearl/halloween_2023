package middleware

import (
	"context"
	"net/http"
	"strings"
)

//Rate limiter
//rps = requests per second
//burst = max requests at once
/*
func (app *app) rateLimit(next http.Handler, rps, burst int) http.Handler {
	clients := make(map[string]*rate.Limiter)

	return http.Handler
}
*/

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Prefer Authorization header
		auth := r.Header.Get("Authorization")
		var tokenStr string
		if strings.HasPrefix(auth, "Bearer ") {
			tokenStr = strings.TrimPrefix(auth, "Bearer ")
		} else {
			// 2. Fallback: access token in cookie (if you prefer)
			c, err := r.Cookie("access_token")
			if err == nil {
				tokenStr = c.Value
			}
		}
		if tokenStr == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
