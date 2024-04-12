package middleware

import (
	"context"
	"log"
	"net/http"
	"time"
)

func TimeOutMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), time.Second)

		r = r.WithContext(ctx)

		defer cancel()
		// replacing request context with new context
		r = r.WithContext(ctx)

		done := make(chan struct{})

		go func() {

			defer close(done)

			next.ServeHTTP(w, r)
		}()

		select {
		case <-r.Context().Done():
			log.Println("Request timeout exceed")
			return

		case <-done:
			return
		}

	})
}
