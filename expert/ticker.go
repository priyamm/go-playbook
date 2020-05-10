package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Ticker stopped")
				return
			case <-ticker.C:
				fmt.Println("Ticker executed")
			}
		}
	}()

	fmt.Println("Ticker running")
	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	done <- true
	fmt.Println("Exiting main")
	time.Sleep(time.Second)
}
