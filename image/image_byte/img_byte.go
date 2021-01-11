package main

import (
        "fmt"
        "io/ioutil"
        "log"
)

func main() {
        content, err := ioutil.ReadFile("image_1.jpg")
        if err != nil{
                log.Fatal(err)
        }

        fmt.Printf("File contents: %s", content)
}
