package main

import "fmt"

func main() {

	// use make to create a map
	var idMap1 map[int]string = make(map[int]string)
	idMap2 := map[int]string{1: "Shyam", 2: "Adi"}

	idMap1[1] = "Rahul"
	idMap1[2] = "Ram"

	fmt.Println(idMap1)
	fmt.Println(idMap2)

	_, isPresent := idMap1[3]
	fmt.Println(isPresent)

	delete(idMap2, 1)
	fmt.Println(idMap2)
}
