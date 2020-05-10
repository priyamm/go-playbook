package main

import (
	"fmt"
	"time"
)

func printSeq() {
	for i := 0; i < 3; {
		fmt.Println("Inside printSeq")
		i++
	}
}

func main() {
	fmt.Println("Inside Main")
	go printSeq()
	time.Sleep(time.Second)
	fmt.Println("Exiting Main")
}
