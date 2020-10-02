package main

import "fmt"

func main() {
	// create empty array
	var empty [3]string
	if empty[0] == "" {
		fmt.Println("Array has no value")
	}

	// create array with values
	filled := []string{"hello", "world"}

	// assign a value from another array
	empty[0] = filled[1]
	if empty[0] == "world" {
		fmt.Println("I have assigned a new value")
	}

	// assign a new value to old array should not affect other array
	filled[1] = "bye"
	if empty[0] == "world" {
		fmt.Println("The values are not passed by reference")
	}

	type Object struct{ name string }
	obj := Object{name: "original"}
	// instatiate a new array with make
	linked := make([]*Object, 1)
	// append will automatically extend the array
	linked = append(linked, &obj)
	// new value is no longer at the start
	if linked[1].name == obj.name {
		fmt.Println("I am assigned to the end")
	}
	// set new value
	obj.name = "new"
	if linked[1].name == "new" {
		fmt.Println("Array value has changed becuase we are referencing a pointer")
	}
}
