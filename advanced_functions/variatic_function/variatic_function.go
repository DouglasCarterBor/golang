package main

import "fmt"


func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	return total

}

func toWrite(text string, numbers ...int) {
	fmt.Println(text)
	for _, num := range numbers {
		fmt.Println(text, num)
	}
}

func main() {
	total := sum(1, 2, 3, 4, 5)
	total1 := sum()
	fmt.Println(total)
	fmt.Println(total1)

	toWrite("Hello", 1, 2, 3, 4, 5)
}
