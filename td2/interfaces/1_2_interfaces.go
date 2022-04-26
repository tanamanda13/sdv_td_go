package main

import (
	"fmt"
	"log"

	// "os"
	"time"
)

func main2() {
	err := run()
	if err != nil {
		log.Fatal(err)
		// os.Exit(0)
	}

	// fmt.Println(err)
}

func (err MyError)Error() string {
	return fmt.Sprintf("Date : %s. Infos : %s", err.When.String(), err.What)
}

func run() error {
	return MyError{When: time.Now(), What: "Erreur"}
}

type MyError struct{
  When time.Time
  What string
}