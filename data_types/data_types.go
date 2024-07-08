package main

import (
	"errors"
	"fmt"
)

func main () {
	int8, int16, int32, int64 := 1, 2, 3, 4
	fmt.Println(int8, int16, int32, int64)

	intVar := 1
	fmt.Println(intVar)

	uint8, uint16, uint32, uint64 := uint8(1), uint16(2), uint32(3), uint64(4)
	fmt.Println(uint8, uint16, uint32, uint64)

	// alias
	// INT32 = RUNE
	var number rune = 123
	fmt.Println(number)

	// alias
	// UINT8 = BYTE
	var number2 byte = 123
	fmt.Println(number2)

	float32Var, float64Var := float32(1.1), float64(1.1)
	fmt.Println(float32Var, float64Var)

    var realNumber float32 = 1.1555
    fmt.Println(realNumber)

	realNumber2 := 1.1555
	fmt.Println(realNumber2)

	var str string = "This is a string"
	fmt.Println(str)

	str2 := "This is another string"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var text string
	fmt.Println(text)

	var boolean bool = true
	fmt.Println(boolean)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}