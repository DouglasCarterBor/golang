package main

import "fmt"

func fibonacci(position uint) uint {
	//! beware of stack overflow
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	fmt.Println("Recursive Functions")

	position := uint(30)
	for i := uint(0); i <= position; i++ {
		fmt.Printf("Fibonacci number at position %d is %d\n", i, fibonacci(i))
	}
	fmt.Printf("Fibonacci number at position %d is %d\n", position, fibonacci(position))
} 
