package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool, 1)
	go func() {
		fmt.Println("Within anonymous")
		time.Sleep(time.Second)
		channel <- true
	}()

	fmt.Println("Within main")
	<-channel
	fmt.Println("Exiting main")
}
