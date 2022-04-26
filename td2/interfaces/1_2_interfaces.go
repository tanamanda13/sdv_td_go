package main

import (
	"fmt"
	"log"

	// "os"
	"time"
)

func main2() {
	err := run() // stocker le retour
	if err != nil {
		log.Fatal(err)
		// os.Exit(0)
	}

	// fmt.Println(err)

	// Correction 
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func (err *MyError)Error() string {
	return fmt.Sprintf("Date : %s. Infos : %s", err.When.String(), err.What)
}

func run() error { // renvoye une erreur
	// return &MyError{When: time.Now(), What: "Erreur"}
	return &MyError{time.Now(), "Erreur"}
}

type MyError struct{
  When time.Time
  What string
}