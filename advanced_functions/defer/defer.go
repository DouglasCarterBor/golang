package main

import "fmt"

func func1() {
	fmt.Println("func1")
}

func func2() {
	fmt.Println("func2")
}

func studentAproval(n1, n2 float32) bool {
	defer fmt.Println("End")
	fmt.Println("Start")
	media := (n1 + n2) / 2
	if media >= 6 {
		fmt.Println("Aproved")
		return true
	}

	fmt.Println("Disapproved")
	return false
}

func main() {
	defer func1()
	func2()

	fmt.Println(studentAproval(7, 8))
}
