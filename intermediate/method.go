/*
Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it. 
With the help of the receiver argument, the method can access the properties of the receiver. 
Here, the receiver can be of struct type or non-struct type. When you create a method in your code the receiver and receiver type must present in the same package. 
And you are not allowed to create a method in which the receiver type is already defined in another package including inbuilt type like int, string, etc. 
If you try to do so, then the compiler will give an error.
*/

package main

import (
	"fmt"
)

type hello struct {
	name string
}

func (i hello) knock() string {
	reply:= fmt.Sprintf(`
	🚪 Knock, knock
	🐼 Who’s there?
	🦄 %[1]s
	🐼 %[1]s who?
	🦄 %[1]s from your go method
	`, i.name)
	
	return reply
}

func main(){
	hi :=  hello{name: "Gopher"}
	fmt.Println(hi.knock())
}