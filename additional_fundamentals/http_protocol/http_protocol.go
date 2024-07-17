package main

import (
	"log"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func usersHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users Page"))
}

func main() {

	http.HandleFunc("/home", handlerFunc)
	http.HandleFunc("/users", usersHandlerFunc)
	log.Fatal(http.ListenAndServe(":5000", nil))
	
}
