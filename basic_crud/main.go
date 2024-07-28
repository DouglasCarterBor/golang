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
	router.HandleFunc("/users", server.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", server.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")

	fmt.Println("Server running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}