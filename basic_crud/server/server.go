package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := io.ReadAll(r.Body)
	if error != nil {
		w.Write([]byte("Error reading body request"))
		return
	}

	var user user
	json.Unmarshal(bodyRequest, &user)

	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		w.Write([]byte("Error converting body request to user struct"))
		return
	}
	

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("insert into users (name, email) values (?, ?)")
	if error != nil {
		w.Write([]byte("Error creating statement"))
		return
	}

	defer statement.Close()

	insertion, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		w.Write([]byte("Error inserting user"))
		return
	}

	id, error := insertion.LastInsertId()
	if error != nil {
		w.Write([]byte("Error getting last id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created successfully with id %d", id)))
	user.ID = int(id)

}
