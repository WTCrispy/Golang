package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type User struct {
	UserName string
	Password string
}

func main() {
	var (
		users       []User
		outputUsers = make(map[string]User)
	)
	data, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &users)
	for index, user := range users {
		i := fmt.Sprintf("%v", index)

		outputUsers["user-"+i] = user
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(outputUsers)

}
