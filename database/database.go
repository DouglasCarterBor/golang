package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "golang:golang@tcp(172.17.144.1:3306)/golang?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Erro ao pingar o banco de dados: ", err)
	}

	fmt.Println("Conexão com o banco de dados está ativa!")

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Erro ao executar a query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("Erro ao ler os dados da row: ", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Erro após leitura das rows: ", err)
	}
}
