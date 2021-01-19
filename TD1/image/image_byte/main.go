package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("image_1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.New()
	h.Write([]byte(content))
	fmt.Printf("%x\n", h.Sum(nil))
}
