package main

import (
	"fmt"
	"sync"
)

func main() {
	var count uint64
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 1; c <= 1000; c++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter value ->", count)

}
