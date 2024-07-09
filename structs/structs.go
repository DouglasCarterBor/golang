package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {

    fmt.Println("Structs in Go")
	// create a new user
	u := user{"John", 30}
	fmt.Println(u)
	fmt.Println(u.name)
	fmt.Println(u.age)

	var u2 user
	u2.name = "Jane"
	u2.age = 25
	fmt.Println(u2)

	u3 := user{name: "Doe", age: 40}
	fmt.Println(u3)

	u4 := new(user)
	u4.name = "Alice"
	u4.age = 35
	fmt.Println(u4)

	u5 := &user{name: "Bob", age: 45}
	fmt.Println(u5)
	
}