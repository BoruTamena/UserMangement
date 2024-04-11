package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/BoruTamena/UserManagement/services"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// getting token
		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, err := services.ParseAccessToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		handler.ServeHTTP(w, r)
	}

}
