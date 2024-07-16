package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age uint `json:"age"`
}

func main() {
	c := dog{
	"Rex", "Bulldog", 3,
	}
	fmt.Println(c)

	dogInJson, error := json.Marshal(c)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(dogInJson)
	fmt.Println(bytes.NewBuffer(dogInJson))

	c2 := map[string]string{
		"name": "Toby",
		"race": "Golden Retriever",
		"age": "5",
	}

	dogInJson2, error := json.Marshal(c2)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(dogInJson2)
	fmt.Println(bytes.NewBuffer(dogInJson2))

}
