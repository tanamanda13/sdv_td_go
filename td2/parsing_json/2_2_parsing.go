package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var usersMap map[string]User // Déclaration + type

func main() {

	file, _ := ioutil.ReadFile("./users.json")
	var users []User
	json.Unmarshal(file, &users)
	fmt.Println(users) // [{matm 123456 5} {fake44 azerty 13}]
	// si on change les majuscule du json, ca ne change rien, contrairement aux var de Golang
	

	usersMap = make(map[string]User) // Affectation
	for i := 0; i < len(users); i++ {
		usersMap[users[i].UserID] = users[i]
	}
	fmt.Println(usersMap)

	http.HandleFunc("/", ServeHTTP)
	log.Fatal(http.ListenAndServe(":4000", nil))
}

type User struct {
	Login string `json:"userName"`
	Password string
	UserID string
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world and people!\n")
}