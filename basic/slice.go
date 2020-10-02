/*
Slice is dynamically-sized.
Can be created with var, slice literal or make built-in function.
Syntax: []T, where T is a type of the elements in the slice.
slice[startIndex:endIndex], where endIndex is not included.

Slice has a length and capacity.
Length is the number of elements it contains.
Capacity is a number of elements in the underlying array, but counting from the first element in the slice.

To add element to the slice, use build-in function append.

To iterate through the slice, you can use keyword `range`.
*/
package main

import (
	"fmt"
)

func main() {
	arrayA := [5]int{1, 2, 3, 4, 5}

	var sliceA []int = arrayA[1:4]
	// 2, 3, 4
	fmt.Println(sliceA)

	sliceLiteral := []bool{true, false, true}
	// true, false, true
	fmt.Println(sliceLiteral)

	sliceB := []int{6, 7, 8, 9, 10, 11}
	// Length of sliceB: 6, capacity of sliceB: 6
	fmt.Printf("Length of sliceB: %d, capacity of sliceB: %d\n", len(sliceB), cap(sliceB))

	// modify the sliceB
	sliceB = sliceB[:4]
	// Now length of sliceB: 4, capacity of sliceB: 6
	fmt.Printf("Now length of sliceB: %d, capacity of sliceB: %d\n", len(sliceB), cap(sliceB))

	sliceC := make([]string, 5)
	fmt.Println(sliceC)

	multiDimensionalSlice := [][]string{
		[]string{"Bob", "27", "star"},
		[]string{"Jake", "35", "water"},
		[]string{"Julia", "45", "fire"},
	}
	// [[Bob 27 star] [Jake 35 water] [Julia 45 fire]]
	fmt.Println(multiDimensionalSlice)

	sliceD := []int{1, 2, 3}
	sliceD = append(sliceD, 4)
	// 1, 2, 3, 4
	fmt.Println(sliceD)

	for idx, value := range sliceD {
		fmt.Printf("Index %d, value %d\n", idx, value)
	}

	// reverse a slice
	for i, j := 0, len(sliceD)-1; i < j; i, j = i+1, j-1 {
		sliceD[i], sliceD[j] = sliceD[j], sliceD[i]
	}
	fmt.Println("Reversed sliceD: ", sliceD)
}
