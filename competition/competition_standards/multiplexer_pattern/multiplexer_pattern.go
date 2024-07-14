package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	channel := multiplexar(write("Hello"), write("World"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}	

}

func multiplexar(inputChannel, inputChannel1 <-chan string) <-chan string {
	outChannel := make(chan string)

	go func() {
		for {
			select {
				case message := <-inputChannel:
					outChannel <- message
				case message := <-inputChannel1:
					outChannel <- message
			}
		}
	} ()

	return outChannel
}

func write (text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Multiplexer: %s", text)
            time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	} ()

	return channel
}