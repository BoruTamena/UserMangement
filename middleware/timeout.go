package middleware

import (
	"context"
	"net/http"
	"time"
)

func TimeOutMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), time.Second)

		r = r.WithContext(ctx)

		defer cancel()
		// replacing request context with new context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
