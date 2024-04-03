package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BoruTamena/UserManagement/db"
	error_code "github.com/BoruTamena/UserManagement/entity"
	"github.com/BoruTamena/UserManagement/models"
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

	if r.Method != http.MethodPost {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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

	if r.Method != http.MethodGet {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// fetching users list
	res_data := hr.Select()

	if res_data.Code == error_code.Success {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(res_data)

}

func (hr handRepo) UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "Method not allowed ", http.StatusMethodNotAllowed)
		log.Print(r.Method)
		return

	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique filename to avoid collisions
	filename := handler.Filename
	uploadDir := "./upload"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, 0755)
	}
	filepath := filepath.Join(uploadDir, filename)
	f, err := os.Create(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Buffered copy for better performance
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Image uploaded successfully!")
}

func password_hashing(password string) string {

	hash_pas, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(hash_pas)

}

func ComparePassword(hash_pass, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash_pass), []byte(password))

	return err == nil

}
