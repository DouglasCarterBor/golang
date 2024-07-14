package main

import (
	"fmt"
	"time"
)

func main() {
	// competition != parallelism
	go write("Hello") // goroutine
	write("World")

}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
