package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	UserId   string
	UserName string
	Password string
}

// GLOBAL VAR
var usersMap = make(map[string]User)

func getJson(u *[]User, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &u)
}
func getUserById(res http.ResponseWriter, req *http.Request) {
	// query id -> ?id=something
	id := req.FormValue("id")
	if id != "" {
		fmt.Printf("Your id: %v\n", id)
		if findelem, ok := usersMap[id]; ok {

			// Set my response to a Json format
			res.Header().Set(
				"Content-Type",
				"application/json;charset=utf-8",
			)
			// return response 200 -> OK
			res.WriteHeader(http.StatusOK)

			// Convert the data to a []byte
			json_data, err := json.Marshal(findelem)
			if err != nil {
				log.Fatal(err)
			}
			// return my data
			_, err = res.Write(json_data)
			return
		}
		// 	return 404  -> not found
		res.WriteHeader(http.StatusNotFound)
	}
}

//////////////////////////////
//MAIN
/////////////////////////////
func main() {
	var (
		users []User
	)
	getJson(&users, "./users.json")
	for _, user := range users {
		usersMap[user.UserId] = user
	}
	http.HandleFunc("/", getUserById)
	http.ListenAndServe(":80", nil)
}
