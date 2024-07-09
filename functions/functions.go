package main

import "fmt"

func sum(n1 int8, n2 int8) int {
	return int(n1) + int(n2)
}

func calc(n1, n2 int8) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return int(sum), int(sub)
}

func main() {
	sm := sum(1, 2)
    fmt.Println("Sum:", sm)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Hello World!")
	fmt.Println(result)

	resultSum, resultSub := calc(10, 5)
	fmt.Println("Sum:", resultSum, "Sub:", resultSub)

	resultSum2, _ := calc(10, 5)
	fmt.Println("Sum:", resultSum2)
}

