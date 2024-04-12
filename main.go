package main

import (
	"log"
	"net/http"

	"github.com/BoruTamena/UserManagement/db"
	"github.com/BoruTamena/UserManagement/handlers"
	"github.com/BoruTamena/UserManagement/middleware"
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

	// creating server mux

	mux := http.NewServeMux()

	mux.HandleFunc("/user", hr.Register)
	mux.HandleFunc("/users", middleware.AddValueMiddleWare(middleware.Auth(hr.ListUser)))
	mux.HandleFunc("/upload", hr.UploadImage)
	mux.HandleFunc("/login", auth.Login)
	mux.HandleFunc("/refresh", auth.Refersh)

	log.Fatal(http.ListenAndServe(":3000", middleware.TimeOutMiddleware(mux)))

}
