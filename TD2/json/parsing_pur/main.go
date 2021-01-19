package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	group := User{
		Login:    "Paul",
		Password: "Pass123",
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))
}

type User struct {
	Login    string `json:"userName"`
	Password string
}
