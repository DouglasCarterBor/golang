package main

import "fmt"

func main() {
    fmt.Println("Maps")

	usuario := map[string]string{
		"nome": "Douglas",
		"sobrenome": "Bordignon",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	user := map[string]map[string]string{	
		"nome": {
			"first": "Douglas",
			"last": "Bordignon",
		},
		"curso": {
			"nome": "Engenharia de Software",
		},
	}
	
	fmt.Println(user)

	delete(user, "nome")
	fmt.Println(user)

	user["age"] = map[string]string{
		"age": "25",
	}
	
	fmt.Println(user)

}