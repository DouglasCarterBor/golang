package main

import "fmt"

type person struct {
	name string
	age int
	surname string
	height float32
}

type student struct {
	person
	course string
	faculty string
}

func main() {
	fmt.Println("Heritage in Go")
	p := person{name: "John", age: 25, surname: "Doe", height: 1.75}

	fmt.Println(p)

	s := student{person: person{name: "Jane", age: 20, surname: "Doe", height: 1.65}, course: "Computer Science", faculty: "Engineering"}

	fmt.Println(s)
    s.age = 21
	s.height = 1.70
	fmt.Println(s.age)
	fmt.Println(s.height)

}