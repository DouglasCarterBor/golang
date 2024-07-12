package main

import "fmt"

func main()  {
	ret := func(text string) string {
		fmt.Println("Anonymous function", text)
		return fmt.Sprintf("Anonymous function -> %s", text)
	}("Passing text to anonymous function")

	fmt.Println(ret)
}
