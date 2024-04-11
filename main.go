package main

import (
	"log"
	"net/http"
	"time"

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

	mux.HandleFunc("/user", middleware.TimeOutMiddleware(hr.Register))
	mux.HandleFunc("/users", middleware.AddValueMiddleWare(middleware.Auth(hr.ListUser)))
	mux.HandleFunc("/upload", handlers.UploadImage)
	mux.HandleFunc("/login", auth.Login)
	mux.HandleFunc("/refresh", auth.Refersh)

	server := &http.Server{

		Addr:         ":3000",
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Handler:      mux,
	}

	log.Fatal(server.ListenAndServe())

}
