package main

import (
	"fmt"
"encoding/json"
)

func main() {
var jsonBlob = []byte(
`[
{"Login": "Paul","Password": "Pass123"},
{"Login": "Jacques","Password": "Jacquouille123"}
]`)
var users []User
err := json.Unmarshal(jsonBlob, &users)
if err != nil {
fmt.Println("error:", err)
}
fmt.Printf("%+v", users)
}

type User struct {
Login string
Password string
}
