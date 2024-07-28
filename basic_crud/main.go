package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE
	router := mux.NewRouter()
	
	router.HandleFunc("/users", server.CreateUser).Methods("POST")

	fmt.Println("Server running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}