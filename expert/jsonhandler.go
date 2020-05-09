package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// for json parsing the struct should have variable starting with capital letter
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age`
}

type UserList []User

func main() {
	userFirst := User{Name: "Ravi", Age: 22}
	userSecond := User{Name: "Rohan", Age: 20}
	userList := UserList{userFirst, userSecond}

	// printing the list of users
	fmt.Println(userList)

	// serializing using Marshal
	jsonValue, _ := json.Marshal(userList)
	fmt.Println(string(jsonValue))

	// deserializing using Unmarshal
	dat, err := ioutil.ReadFile("../resources/userlist.json")
	if err != nil {
		fmt.Println(err)
	}
	var user UserList
	json.Unmarshal(dat, &user)
	fmt.Println("User present in file -> ", user)

}
