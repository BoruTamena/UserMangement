package middleware

import (
	"context"
	"errors"
	"net/http"

	error_code "github.com/BoruTamena/UserManagement/entity"
	"github.com/BoruTamena/UserManagement/handlers"
	"github.com/BoruTamena/UserManagement/services"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// getting token
		token := r.Header.Get("Authorization")

		if token == "" {
			erobj := handlers.NewError(errors.New("Auth token is not provide"), handlers.ErrorT(error_code.Unauthorized), 500)
			ctx := context.WithValue(r.Context(), "err", erobj)
			r = r.WithContext(ctx)
			erobj.HandleError(w, r)
			return

		}

		_, err := services.ParseAccessToken(token)

		if err != nil {
			erobj := handlers.NewError(errors.New("token"), handlers.ErrorT(error_code.Unauthorized), 500)
			ctx := context.WithValue(r.Context(), "err", erobj)
			r = r.WithContext(ctx)
			erobj.HandleError(w, r)
			return

		}

		handler.ServeHTTP(w, r)
	}

}
