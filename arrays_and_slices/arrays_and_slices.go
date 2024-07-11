package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("arrays_and_slices")

	var array [5]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	array1 := [5]string{"1 position", "2 position", "3 position", "4 position", "5 position"}
	fmt.Println(array1)

	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := array1[1:3]
	fmt.Println(slice2)

	array1[1] = "changed position"
	fmt.Println(slice2)

}