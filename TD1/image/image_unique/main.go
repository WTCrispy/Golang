package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func hashImg(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.New()
	h.Write(content)
	return h.Sum(nil)
}
func main() {

	image1 := hashImg("image_1.jpg")
	image2 := hashImg("image_2.jpg")
	image3 := hashImg("image_3.jpg")

	compare1 := bytes.Compare(image1, image2)
	compare2 := bytes.Compare(image1, image3)
	compare3 := bytes.Compare(image2, image3)

	if compare1 != 0 && compare2 != 0 {
		fmt.Println("image_1 est unique")
	} else {
		fmt.Println("image_1 n'est pas unique")
	}

	if compare1 != 0 && compare3 != 0 {
		fmt.Println("image_2 est unique")
	} else {
		fmt.Println("image_2 n'est pas unique")
	}

	if compare2 != 0 && compare3 != 0 {
		fmt.Println("image_3 est unique")
	} else {
		fmt.Println("image_3 n'est pas unique")
	}

}
