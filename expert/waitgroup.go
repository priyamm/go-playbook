package main

import (
	"fmt"
	"sync"
	"time"
)

func heavyTask(index int, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("Index is -> ", index)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go heavyTask(i, &wg)
	}

	fmt.Println("Processing main")
	wg.Wait()
	fmt.Println("Exiting main")
}
