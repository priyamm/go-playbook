package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Take input from user and add to all the previous inputs and sort it, exits only when user press 'X'

func main() {
	sli := make([]int, 0, 3)
	var input string
	for {
		fmt.Println("Enter Integer value or press X to exit")
		_, _ = fmt.Scan(&input)
		if "X" == input {
			break
		}
		str, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Not a number")
		} else {
			sli = append(sli, str)
			sort.Ints(sli)
			fmt.Println(sli)
		}
	}
}
