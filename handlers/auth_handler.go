package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BoruTamena/UserManagement/db"
	"github.com/BoruTamena/UserManagement/models"
	"github.com/BoruTamena/UserManagement/services"
)

type authRepo struct {
	*db.UserDb
}

// constructor
func NewAuthHandler(userdb *db.UserDb) *authRepo {
	return &authRepo{
		userdb,
	}
}

func (hr authRepo) Login(w http.ResponseWriter, r *http.Request) {

	var user models.UserLogIn

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Print(err)

	}

	// checking user creditialities

	//creating token
	token, err := services.CreateToken(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	w.WriteHeader(http.StatusAccepted)

	log.Print(token)
	json.NewEncoder(w).Encode(token)

}

func Auth(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// performing authentication here

		handler.ServeHTTP(w, r)

	}
}
