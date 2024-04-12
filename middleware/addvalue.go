package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"

	error_code "github.com/BoruTamena/UserManagement/entity"
	"github.com/BoruTamena/UserManagement/handlers"
	"github.com/BoruTamena/UserManagement/services"
	"github.com/golang-jwt/jwt/v5"
)

const (
	ReqId  = "ReqId"
	UserId = "UserId"
)

func generate_rand_id() (string, error) {

	read_byte := make([]byte, 8)

	// reading
	_, err := rand.Read(read_byte)

	if err != nil {

		return "", err
	}

	requestId := hex.EncodeToString(read_byte)

	return requestId, nil

}

func AddValueMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		accesstoken := r.Header.Get("Authorization")

		if accesstoken == "" {

			erobj := handlers.NewError(errors.New("Access token is not provided"), handlers.ErrorT(error_code.Unauthorized), 500)
			ctx := context.WithValue(r.Context(), "err", erobj)
			r = r.WithContext(ctx)
			erobj.HandleError(w, r)
			return

		}

		request_id, err := generate_rand_id()

		if err != nil {

			erobj := handlers.NewError(err, handlers.ErrorT(error_code.UnableToSave), 500)
			ctx := context.WithValue(r.Context(), "err", erobj)
			r = r.WithContext(ctx)
			erobj.HandleError(w, r)
			return

		}

		token, err := services.ParseAccessToken(accesstoken)
		if err != nil {

			erobj := handlers.NewError(errors.New("Invalid access token"), handlers.ErrorT(error_code.Unauthorized), 500)
			ctx := context.WithValue(r.Context(), "err", erobj)
			r = r.WithContext(ctx)
			erobj.HandleError(w, r)
			return

		}

		// Extract user ID from refresh token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {

			return
		}
		userID, ok := claims["userId"].(float64)

		// adding request id
		ctx := context.WithValue(r.Context(), ReqId, request_id)
		// adding user_id
		ctx = context.WithValue(ctx, UserId, int(userID))

		r = r.WithContext(ctx) // replace request context

		next.ServeHTTP(w, r)
	}
}
