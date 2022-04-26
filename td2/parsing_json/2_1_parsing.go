package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// user1 := User {
	// 	Login: "Paul",
	// 	Password: "pass123",
	// }
	// b, err := json.Marshal(user1)
	// if err != nil {
  //   fmt.Println("error:", err)
  // }
	// fmt.Println(string(b))


	// // Open our jsonFile
	// jsonFile, err := os.Open("./users.json")
	// // if we os.Open returns an error then handle it
	// if err != nil {
	// 		fmt.Println(err)
	// }
	// fmt.Println("Successfully Opened users.json")
	// // defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	jsonFile, err := os.Open("./users.json")
	if err != nil {
		fmt.Println(err)
	}
	var users []User
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &users)

	// fmt.Println(users)
	// usersList := make([]string,len(users))
	usersList := make([]User,len(users))
	for i, u := range users {
		usersList[i] = u
		// fmt.Println(i)
		// fmt.Println(u) // users[i]
	}
	fmt.Println(usersList)
}

type User struct {
	Login string `json:"userName"`
	Password string
}