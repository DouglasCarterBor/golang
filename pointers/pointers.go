package main

import "fmt"

func main() {
	fmt.Println("pointers")

	var variable int = 10
	var variable1 int = variable

	fmt.Println(variable, variable1)
	variable++

	var variable2 int
	var pointer *int

    variable2 = 100
	pointer = &variable2

	fmt.Println(variable2, pointer)

	variable2 = 150
	fmt.Println(variable2, pointer)
	fmt.Println(variable2, *pointer)

}