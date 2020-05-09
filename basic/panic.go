package main

import "time"

func main() {
	//  panic is used to represent any unwanted or unexpeected error and we don't know or don't want to handle
	//panic("Something went wrong!")

	go asyncMethod()
	time.Sleep(time.Second)
}

func asyncMethod() {
	panic("panic in asyncMethod")
}
