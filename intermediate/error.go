package main

import (
	"errors"
	"fmt"
)

func errorSupport() (int, int, int, error) {
	return 1, 2, 3, errors.New("Error Thrown")
}

func main() {
	i, j, k, err := errorSupport()
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(err)
}
