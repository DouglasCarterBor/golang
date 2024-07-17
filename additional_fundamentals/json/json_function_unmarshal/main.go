package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age uint `json:"age"`
}

func main () {
	dogInJSON := `{"name":"Firulais", "race":"Pitbull", "age":3}`
	var dog1 dog
	if error := json.Unmarshal([]byte(dogInJSON), &dog1); error != nil {
		log.Fatalf("Error unmarshalling the JSON: %v", error)
	}
	fmt.Printf("Dog: %v\n", dog1)

	dogInJSON2 := `{"name":"Toby", "race":"Poodle", "age":5}`

	c2 := make(map[string]interface{})
	if err := json.Unmarshal([]byte(dogInJSON2), &c2); err != nil {
		log.Fatalf("Error unmarshalling the JSON: %v", err)
	}

	fmt.Printf("Dog: %v\n", c2)

}
