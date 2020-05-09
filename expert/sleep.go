package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Index is -> ", i)
		time.Sleep(time.Second)
	}
}
