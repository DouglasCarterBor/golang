package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go writeChannel("Hello", channel)

	fmt.Println("After writeChannel")

	// for {
	// 	message, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("End of main")

}

func writeChannel(text string, channel chan string) {

	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(1 * time.Second)
	}

	close(channel)

}
