package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main() {
	wg.Add(1)
	go maFonction(wg)
	fmt.Println("Fin du programme")	
}

func maFonction(wg sync.WaitGroup) {
	fmt.Println("j'ai fini !")
	defer wg.Done()
}