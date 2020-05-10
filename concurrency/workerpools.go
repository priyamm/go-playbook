package main

import (
	"fmt"
	"time"
)

func heavyTask(id int, num int, result chan<- int) {
	fmt.Println("Worker", id, "Executing task", num)
	time.Sleep(time.Second)
	result <- num
	fmt.Println("Worker", id, "Completed task", num)
}

func main() {
	jobs := make(chan int, 5)
	result := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go func(id int, result chan<- int) {
			for j := range jobs {
				heavyTask(id, j, result)
			}
		}(i, result)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// have to solve error for state issue
	for _ = range result {
		fmt.Println(<-result)
	}
}
