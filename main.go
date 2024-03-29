package main

import (
	"log"
	"net/http"

	"github.com/BoruTamena/UserManagement/handlers"
)

func main() {

	// starting server

	http.HandleFunc("/user", handlers.Register)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
