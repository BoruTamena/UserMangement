package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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

func (ah authRepo) getuser(username string) models.UserReg {

	data := ah.Data

	for _, item := range data {

		if strings.EqualFold(item.UserName, username) {
			return item
		}
	}

	return models.UserReg{}
}

func (ah authRepo) Login(w http.ResponseWriter, r *http.Request) {

	var user models.UserLogIn

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Print(err)

	}

	// checking user creditialities

	user_detail := ah.getuser(user.UserName)

	if ok := ComparePassword(user_detail.Password, user.Password); !ok {

		w.WriteHeader(http.StatusUnauthorized)

		return
	}

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
