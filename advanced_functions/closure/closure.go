package main

import "fmt"

func closure() func() {
	text := "Inside the closure"

	func1 := func() {
		fmt.Println(text)
	}

	return func1
}

func main() {
	text := "Inside the main function"
	fmt.Println(text)

	newFunc := closure()
	newFunc()

}