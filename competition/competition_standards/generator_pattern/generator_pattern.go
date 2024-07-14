package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("World")
    for i := 0; i < 5; i++ { 
		fmt.Println(<-channel)
	}

}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Hello %s", text)
			time.Sleep(1 * time.Second)
		}
	}()

	return channel
}