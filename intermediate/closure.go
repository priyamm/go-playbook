package main

import (
	"fmt"
)

func increment(msg string) func() {
	i := 0
	return func() {
		i++
		fmt.Printf("Value of i in %s -> %d\n", msg, i)
	}
}

func main() {
	incInstanceFirst := increment("incInstanceFirst")
	incInstanceSecond := increment("incInstanceSecond")

	incInstanceFirst()
	incInstanceFirst()
	incInstanceSecond()
	incInstanceFirst()
	incInstanceSecond()
}
