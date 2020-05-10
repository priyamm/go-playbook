package main

import (
	"fmt"
)

func main() {
	channel := make(chan string) // these are unbuffered channel, goroutine is blocked until the value is received from the channel
	defer(channel) // to close the channel after the work is done so that no goroutine can send anything
	go func() {
		channel <- "Setting msg in Channel"
	}()

	value := <-channel
	fmt.Println("Value Exchanged -> ", value)
}
