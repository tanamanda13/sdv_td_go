package main

import (
	"fmt"
	// "reflect"
)

func main() {
	PrintIt(true)
}

func PrintIt(input interface{}) {
	switch input.(type) { // v := input.(type)
	case int:
		fmt.Println("c'est un int")
	case string:
		fmt.Println("c'est un string")
	default:
		fmt.Println("c'est autre chose d'un string ou int")
	}
}
