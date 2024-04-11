package middleware

import (
	"log"
	"net/http"

	error_code "github.com/BoruTamena/UserManagement/entity"
)

type Error struct {
	ErrorCode int
	ErrorType string
	ErrorMsg  string
}

func ErrorMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// serving the request
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))

		// Reading error from context
		// handlers.HandleError(w, r)

		err := ctx.Value("err")

		log.Println("error form middleware", ctx.Value("err"))

		if err == nil {

			return

		}
		if e, ok := err.(*Error); ok {
			// Handle the error based on its type
			switch e.ErrorType {
			case string(error_code.UnableToSave):
				http.Error(w, e.ErrorMsg, http.StatusInternalServerError)
			case string(error_code.UnableToFindResource):
				http.Error(w, e.ErrorMsg, http.StatusNotFound)
			case string(error_code.UnableToRead):
				http.Error(w, e.ErrorMsg, http.StatusInternalServerError)
			case string(error_code.Unauthorized):
				http.Error(w, e.ErrorMsg, http.StatusUnauthorized)
			default:
				// Default error response
				http.Error(w, e.ErrorMsg, http.StatusInternalServerError)
			}
		}

	}

}
