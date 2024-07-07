package main

import (
	"fmt"
	"modulo/assistant"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	assistant.Write()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}