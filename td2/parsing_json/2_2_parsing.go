package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	// "log"
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

	// http.Handle("/", &User{}) // en tapant sur localhost, si l'objet a bien la méthode ServeHTTP on va chercher l'objet User
	http.HandleFunc("/", ServeHTTP)
	log.Fatal(http.ListenAndServe(":4001", nil)) // Démarre le server et écoute
}

type User struct {
	Login string `json:"userName"`
	Password string
	UserID string
}

// func (s *User) ServeHTTP(w http.ResponseWriter, r *http.Request) { // On rattache notre serverHTTP à un User -> Dispose maintenant de cette méthode. Si l'objet (User) possède cette méthode ServeHTTP, alors il peut répondre. Mon objet User (receiver) devient un Handler.
func ServeHTTP(w http.ResponseWriter, r *http.Request) { // r -> requet (dans barre recherche, question) et w -> response
	// io.WriteString(w, "Hello, world and people!\n")
	isUserID, userIDKey := searchId(r)
	// fmt.Println(userIDKey)

	sendResponse(w, isUserID, userIDKey)
}

func searchId(r *http.Request) (bool, string) {
	id := r.FormValue("id")
	isUserID := false
	userIDKey := ""
	for k := range usersMap {
		if k == id {
			isUserID = true
			userIDKey = k
			break
		}
	}
	return isUserID, userIDKey
}

func sendResponse(w http.ResponseWriter, isUserID bool, userIDKey string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if isUserID {
		w.WriteHeader(http.StatusOK)
		b1, _ := json.Marshal(usersMap[userIDKey])
		w.Write(b1)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}