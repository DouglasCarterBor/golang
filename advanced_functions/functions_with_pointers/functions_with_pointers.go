package main

import "fmt"

func inverterSignal(x int) int {
   return x * -1
}

func inverterSignalWithPointer(x *int) {
	*x = *x * -1
}

func main() {
    number := 20
	invertedNumber := inverterSignal(number)
	fmt.Println("Before inverterSignal:", number)
	fmt.Println("After inverterSignal:", invertedNumber)

	newNumber := 30
	fmt.Println("Before inverterSignalWithPointer:", newNumber)
	inverterSignalWithPointer(&newNumber)
	fmt.Println("After inverterSignalWithPointer:", newNumber)

}