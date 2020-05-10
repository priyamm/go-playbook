package main

import (
	"fmt"
	"time"
)

func timerTask() {
	fmt.Println("timerTask is executing")
}

func main1() {
	// timer runs after the designated time
	timer := time.NewTimer(2 * time.Second)

	<-timer.C
	timerTask()

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer worked")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer stopped")
	}
}
