package main

import "fmt"

func main() {
	fmt.Println("Control Structures")

	number := 10
    
	if number > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less than 15")
	}

	if number1 := 10; number1 > 15 {
		fmt.Println("Greater than 15")
	} else {
		fmt.Println("Less than 15")
	}

	if number2 := 10; number2 > 15 {
		fmt.Println("Greater than 15")
	} else if number2 < 5 {
		fmt.Println("Less than 5")
	} else {
		fmt.Println("Between 5 and 15")
	}

}