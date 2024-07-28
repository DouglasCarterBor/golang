package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver Connection to MySQL
)

// Connect - Function to connect to the database
func Connect() (*sql.DB, error) {
	stringConnection := "golang:golang@tcp(172.17.144.1:3306)/golang?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", stringConnection)
	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil

}