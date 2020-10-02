package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func main() {

	fmt.Printf("Dividing %v by %v\n", 4, 2)
	result, err := divide(4, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The division result is %v\n", result)
	}

	fmt.Printf("Dividing %v by %v\n", 4, 0)
	result, err = divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The division result is %v\n", result)
	}

}
