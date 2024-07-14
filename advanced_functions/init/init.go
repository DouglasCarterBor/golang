package main

import "fmt"

var n int

func init() {
	fmt.Println("Init 1")
	n = 10
}

func main() {
	fmt.Println("Func init")
	fmt.Println(n)
}
