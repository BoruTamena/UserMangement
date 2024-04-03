package main

import (
	"log"
	"net/http"

	"github.com/BoruTamena/UserManagement/db"
	"github.com/BoruTamena/UserManagement/handlers"
	"github.com/joho/godotenv"
)

func main() {

	// loading env variable

	err := godotenv.Load(".env")

	if err != nil {
		log.Print(err.Error())

		return
	}

	// creating user db

	user := db.NewUserDb()

	// create repo

	hr := handlers.NewHandler(user)
	auth := handlers.NewAuthHandler(user)

	// starting server
	http.HandleFunc("/user", hr.Register)
	http.HandleFunc("/users", handlers.Auth(hr.ListUser))
	http.HandleFunc("/upload", hr.UploadHandler)
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/refresh", auth.Refersh)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
