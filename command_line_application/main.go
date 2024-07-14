package main

import (
	"command_line_application/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start Point")

	aplication := app.ToGenerate()
	if error := aplication.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}