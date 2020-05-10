package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Setting msg in Channel"
	}()

	value := <-channel
	fmt.Println("Value Exchanged -> ", value)
}
