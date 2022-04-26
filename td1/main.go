package main

import (
	"flag"
	"fmt"
	"os"
)
const VERSION = "1.0"

func main() {
	/*
	var pBool *bool // pointer to Bool (p -> pointer)
	*/

	versionArg := flag.Bool("version", false, "diplay version").
	flag.Parse() // Analyse les argument passés
	/*
	var version *bool
	*version = false
	flag.BoolVar("version", false, "diplay version") // met le résultat dans une var.
	flag.BoolVar("v", false, "diplay version")
	*/
	if *versionArg { // ajouter * permet d'aller voir la valeur du bool, et pas seulement le pointer
		fmt.Println("Version ", VERSION)
		os.Exit(0) // sortir du programme
	}
}
