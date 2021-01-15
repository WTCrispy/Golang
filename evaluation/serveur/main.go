package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {

	h1 := func(rw http.ResponseWriter, _ *http.Request) {
		fmt.Println("/")
		return
	}

	http.HandleFunc("/", h1)

	h2 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("/done")
		return
	}

	http.HandleFunc("/done", h2)

	h3 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("/add")
		return
	}

	http.HandleFunc("/add", h3)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

var task []Task

type Task struct {
	Description string `json:"desc"`

	Done bool `json:"done"`
}

type List struct {
	ID string `json:"id"`
	Task string `json:"task"`
}

func list(rw http.ResponseWriter, _ *http.Request) {
	list := []List{}
	task = []Task{
		{"Faire les courses", false},
		{"Payer les factures", false},
	}
	for id, i := range task {
		if !i.Done {
			list = append(list, List{strconv.Itoa(id), i.Description})
		}
	}
	rw.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(list)

	rw.Write(data)
	return
}
