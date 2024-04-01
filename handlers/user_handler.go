package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/BoruTamena/UserManagement/db"
	error_code "github.com/BoruTamena/UserManagement/entity"
	"github.com/BoruTamena/UserManagement/models"
	"github.com/BoruTamena/UserManagement/services"
	"golang.org/x/crypto/bcrypt"
)

type handRepo struct {
	*db.UserDb
}

func NewHandler(userdb *db.UserDb) *handRepo {
	return &handRepo{
		userdb,
	}
}

func (hr handRepo) Register(w http.ResponseWriter, r *http.Request) {

	var user models.UserReg

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// encrypting password

	user.Password = password_hashing(user.Password)

	// inserting user into
	res_data := hr.Insert(user)

	if res_data.Code == error_code.InvalidRequest {

		w.WriteHeader(http.StatusBadRequest)
	}

	if res_data.Code == error_code.Success {

		w.WriteHeader(http.StatusCreated)

	}

	json.NewEncoder(w).Encode(res_data)

}

func (hr handRepo) ListUser(w http.ResponseWriter, r *http.Request) {

	// getting token
	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := services.ParseAccessToken(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	// fetching users list
	res_data := hr.Select()

	if res_data.Code == error_code.Success {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(res_data)

}

func (hr handRepo) UploadFile(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // parsing

	if err != nil {

		log.Fatal(err)
	}

	file, handler, err := r.FormFile("file") // retirive file

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print(err)

		return
	}

	defer file.Close()

	// creating new file

	newfile, err := os.Create("./upload" + handler.Filename)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer newfile.Close()

	_, err = io.Copy(newfile, file)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode("File Uploaded successfully")

}

func password_hashing(password string) string {

	hash_pas, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(hash_pas)

}

func ComparePassword(hash_pass, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash_pass), []byte(password))

	return err == nil

}
