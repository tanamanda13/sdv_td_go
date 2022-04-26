package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user1 := User {
		Login: "Paul",
		Password: "pass123",
	}
	b, err := json.Marshal(user1)
	if err != nil {
    fmt.Println("error:", err)
  }
	fmt.Println(string(b))
}

type User struct {
	// Login string `json:"login"`
	Login string
	Password string
}