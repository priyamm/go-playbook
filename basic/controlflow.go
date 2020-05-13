package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Index ->", i)
	}

	var j int
	for j = 0; j < 3; j++ {
		fmt.Println("Index ->", j)
	}

	var x int = 22
	if x == 22 {
		fmt.Println("Correct")
	} else {
		fmt.Println("Incorrect")
	}
}
