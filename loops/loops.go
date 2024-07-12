package main

import (
	"fmt"
	//"time"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Incrementando i:", i)
	// 	i++
	// }

	// fmt.Println("Valor final de i:", i)

	// for j := 0; j < 10; j++ {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("Incrementando j:", j)
	// }

	//names := [3]string{"Douglas", "Juan", "Diego"}
	names := []string{"Douglas", "Juan", "Diego"}

	for index, nome := range names {
		fmt.Println(index, nome)
	}
	for _, nome := range names {
		fmt.Println(nome)
	}

	for index, letter := range "word" {
		fmt.Println(index, string(letter))
		fmt.Println(index, letter)
	}

	user := map[string]string {
		"name": "Douglas",
		"lastname": "Bordignon",
	 }

	 for key, value := range user {
		 fmt.Println(key, value)
	 }

	 fmt.Println(user)
	
	 type userStruct struct {
		nome string
		lastName string
	 }

	 //! you can't use a range in a struct
	//  userStruct1 := userStruct{"Douglas", "Bordignon"}

	//  for key, value := range userStruct1 {
	// 	 fmt.Println(key, value)	
	//  }

	//for {
		//! infinite loop
	//}

}