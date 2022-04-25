package main

import (
	"flag"
	"fmt"
)
const VERSION = "1.0"

func main() {
	versionArg := flag.Bool("version", false, "diplay version")
	flag.Parse()
	if *versionArg {
		fmt.Println("Version ", VERSION)
	}
}
