package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var tasks []Task

type Task struct {
	Description string
	Done        bool
}

func (t Task) String() string {
	return fmt.Sprintf("%s", t.Description)
}

func main() {
	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)

	http.ListenAndServe(":8080", nil)
}

func list(rw http.ResponseWriter, _ *http.Request) {
	response := ""

	for i, elt := range tasks {
		if !elt.Done {
			response = fmt.Sprintf("%s\nID:%d, task: %s", response, i, elt.String())
		}
	}

	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte(response))
	if err != nil {
		fmt.Println(err)
	}
}

func done(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // List all done tasks
		res := ""
		for i, elt := range tasks {
			if elt.Done {
				res = fmt.Sprintf("%s\nID: %d task: %s", res, i, elt)
			}
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(res))

	case http.MethodPost: // Set a task to done
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		strData := string(data)
		id, err := strconv.Atoi(strData)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		tasks[id].Done = true
		rw.WriteHeader(http.StatusOK)

	default:
		rw.WriteHeader(http.StatusBadRequest)
	}
}

func add(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		http.Error(
			rw,
			"can't read body", http.StatusBadRequest,
		)
		return
	}

	t := Task{
		Description: string(body),
	}

	tasks = append(tasks, t)

	rw.WriteHeader(http.StatusOK)
}
