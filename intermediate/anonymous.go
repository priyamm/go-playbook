package main

import (
	"fmt"
)

func main() {
	// below is an example of anonymous function
	func(index int) {
		fmt.Println("Index is -> ", index)
	}(1000)

	// we can also pass the reference to a variable
	anonymous := func(str string) int {
		fmt.Println("Input String is -> ", str)

		// len is used for finding the length, its present in the standard library
		return len(str)
	}
	fmt.Println("Address space of anonymous function -> ", anonymous)
	fmt.Println("Length of input string -> ", anonymous("VIRTUAL"))
}
