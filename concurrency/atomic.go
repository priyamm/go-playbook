package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	counter := 0
	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 1; c <= 1000; c++ {
				counter++
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter value ->", counter) // this gives wrong value

	var counter2 uint64
	for i := 1; i <= 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 1; c <= 1000; c++ {
				atomic.AddUint64(&counter2, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter2 value ->", counter2)

}
