package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// used ioutil package
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("../resources/testfile.txt", message, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// used os package
	file, err := os.Create("../resources/testfile2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	numwritten, err := file.Write(message)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Written", numwritten, "bytes")
	}

}
