package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/BoruTamena/UserManagement/db"
	error_code "github.com/BoruTamena/UserManagement/entity"
	"github.com/BoruTamena/UserManagement/validation"

	"github.com/BoruTamena/UserManagement/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	keyerr = "err"
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

	// time.Sleep(time.Second * 2) // simulate

	if r.Method != http.MethodPost {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {

		erobj := newError(err, ErrorT(error_code.UnableToSave), 500)
		ctx := context.WithValue(r.Context(), "err", erobj)
		r = r.WithContext(ctx)
		erobj.HandleError(w, r)
		return
	}

	if !validation.ValidateUser(user) {

		erobj := newError(err, ErrorT(error_code.UnableToSave), 500)
		ctx := context.WithValue(r.Context(), "err", erobj)
		r = r.WithContext(ctx)
		erobj.HandleError(w, r)
		return
	}

	// encrypting password
	user.Password = password_hashing(user.Password)

	// inserting user into
	res_data := hr.Insert(user)

	json.NewEncoder(w).Encode(res_data)

}

func (hr handRepo) ListUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if page := r.URL.Query().Get("page"); page != "" {

		hr.ListUserPagination(w, r)

		return

	} else {
		// fetching users list
		res_data := hr.Select()

		// if res_data.Code == error_code.Success {
		// 	w.WriteHeader(http.StatusOK)
		// }

		json.NewEncoder(w).Encode(res_data)

	}

}

func (hr handRepo) ListUserPagination(w http.ResponseWriter, r *http.Request) {

	pagination := models.PaginationReq{
		Page:     r.URL.Query().Get("page"),
		PageSize: r.URL.Query().Get("page_size"),
	}

	page, err := strconv.Atoi(pagination.Page)

	if err != nil {

		erobj := newError(err, ErrorT(error_code.UnableToSave), 500)
		ctx := context.WithValue(r.Context(), "err", erobj)
		r = r.WithContext(ctx)
		erobj.HandleError(w, r)
		return

	}

	pagesize, err := strconv.Atoi(pagination.PageSize)

	if err != nil {

		erobj := newError(err, "Un Able To Read ", 500)
		ctx := context.WithValue(r.Context(), "err", erobj)
		r = r.WithContext(ctx)
		erobj.HandleError(w, r)
		return

	}

	offset := (page - 1) * pagesize

	// fetching data

	data := hr.SelectPagination(pagesize, offset)

	meta_data := models.MetaData{
		Page:    page,
		PerPage: pagesize,
	}

	res_data := hr.CreateResponse(meta_data, data)

	json.NewEncoder(w).Encode(res_data)

}

func (hr handRepo) UploadImage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("parse err:", err)
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("read error:", err)
		return
	}
	defer file.Close()

	// Save the file
	fileName := handler.Filename
	if fileName == "" {
		http.Error(w, "Image file name is empty", http.StatusBadRequest)
		log.Println("read file name:", err)
		return
	}

	f, err := os.OpenFile(filepath.Join("./uploads", fileName), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Open file err:", err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	// save to resource

	user_id := r.Context().Value("UserId")
	for i, val := range hr.Data {

		if val.Id == user_id {

			hr.Data[i].Image = append(hr.Data[i].Image, f.Name())
			log.Print("user", val)
			break
		}

	}

	w.WriteHeader(http.StatusCreated)
}

func (hr handRepo) GetImage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		return
	}

	user_id := r.Context().Value("UserId")

	var user_img []string

	for i, val := range hr.Data {

		if val.Id == user_id {
			user_img = hr.Data[i].Image
			break
		}
	}

	res_data := map[string]any{
		"img": user_img,
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(res_data)

}

func (hr handRepo) CreateResponse(metadata models.MetaData, Data interface{}) *models.ResponseData {
	return &models.ResponseData{
		Metadata: metadata,
		Data:     Data,
	}
}

func password_hashing(password string) string {

	hash_pas, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	return string(hash_pas)

}

func ComparePassword(hash_pass, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash_pass), []byte(password))

	return err == nil

}
