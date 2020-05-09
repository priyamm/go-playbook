package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// used ioutil package
	dat, err := ioutil.ReadFile("../resources/user.json")
	if err != nil {
		panic(err)
	}
	// dat is basically a []byte
	fmt.Println(string(dat))

	// used os package
	file, err := os.Open("../resources/user.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteReader := make([]byte, 100)
	file.Read(byteReader)
	fmt.Println(string(byteReader))

}
