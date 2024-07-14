package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic(1)
	generic("string")
	generic(1.0)
	generic([]int{1, 2, 3})
	generic(map[string]int{"key": 1})
	generic(struct{ Name string }{"name"})

	fmt.Println(1, 2, "string", 1.0, []int{1, 2, 3}, map[string]int{"key": 1}, struct{ Name string }{"name"})

	//! Example od what not to do
	mapa := map[interface{}]interface{} {
		1: "one",
		float32(2.0): "two",
		"three": 3,
		4: "four",
	}

	fmt.Println(mapa)
}
