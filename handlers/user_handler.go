package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/BoruTamena/UserManagement/db"
	error_code "github.com/BoruTamena/UserManagement/entity"
	"github.com/BoruTamena/UserManagement/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.UserReg

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// creating a database
	user_db := db.UserDb{}

	// inserting user into
	res_data := user_db.Insert(user)

	if res_data.Code == error_code.InvalidRequest {

		w.WriteHeader(http.StatusBadRequest)
	}

	if res_data.Code == error_code.Success {

		w.WriteHeader(http.StatusCreated)

	}

	json.NewEncoder(w).Encode(res_data)

}
