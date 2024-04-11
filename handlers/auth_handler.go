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

func (ah authRepo) getuser(username string) models.User {

	data := ah.Data

	for _, item := range data {

		if strings.EqualFold(item.UserName, username) {
			return item
		}
	}

	return models.User{}
}

func (ah authRepo) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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
	token, refresh, err := services.CreateToken(user_detail)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Print(err)
	}

	resp := map[string]string{

		"access_token":  token,
		"refresh_token": refresh,
	}
	w.WriteHeader(http.StatusAccepted)

	log.Print(token)
	json.NewEncoder(w).Encode(resp)

}

func (ah authRepo) Refersh(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	refToken := models.RefreshTokenReq{}

	err := json.NewDecoder(r.Body).Decode(&refToken)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validating the token

	_, err = services.ParseRefreshToken(refToken.RefreshToken)

	if err != nil {

		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	// generating new token

	accesstoken, err := services.RefreshAccessToken(refToken.RefreshToken)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	res := map[string]string{

		"access": accesstoken,
	}

	json.NewEncoder(w).Encode(res)

}
