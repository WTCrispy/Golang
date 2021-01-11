package main

import (
	"flag"
	"fmt"
)

const VERSION = "1.0"

func main() {
	var version = flag.Bool("version", true, "Version prog")
	fmt.Println("Hello World")
	if *version == true {
		fmt.Println(VERSION)
	}
}
