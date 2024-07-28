package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}
	defer db.Close()

	rows, error := db.Query("select * from users")
	if error != nil {
		w.Write([]byte("Error fetching users"))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user
		rows.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, error := strconv.ParseUint(params["id"], 10, 64)
	if error != nil {
		w.Write([]byte("Error converting id to int"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}

	row, error := db.Query("select * from users where id = ?", ID)
	if error != nil {
		w.Write([]byte("Error fetching user"))
		return
	}

	var user user
	if row.Next() {
		if error := row.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			w.Write([]byte("Error scanning user"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Error encoding user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
