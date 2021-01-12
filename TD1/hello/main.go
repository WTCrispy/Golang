package main

import (
	"flag"
	"fmt"
)

const VERSION = "1.0"

func main() {
	var version = flag.Bool("version", false, "Version prog")
	flag.Parse()
	if *version {
		fmt.Println(VERSION)
	} else {
		fmt.Println("Hello World")
	}
}
