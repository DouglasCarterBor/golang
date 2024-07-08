package main

import "fmt"

func main() {
	var variable string = "This is a variable"
	variable2 := "This is another variable 2"
    fmt.Println(variable)
	fmt.Println(variable2)

	var (
		variable3 string = "This is a variable 3"
		variable4 string = "This is a variable 4"
	)

	fmt.Println(variable3)
	fmt.Println(variable4)

	variable5, variable6 := "This is a variable 5", "This is a variable 6"
	fmt.Println(variable5)
	fmt.Println(variable6)

	const constant string = "This is a constant"
	fmt.Println(constant)

	variable5, variable6 = variable6, variable5
	fmt.Println(variable5)
	fmt.Println(variable6)
}