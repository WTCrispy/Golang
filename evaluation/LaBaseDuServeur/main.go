//Créez un nouveau projet go :
package main

import (
	"fmt"
	"log"
	"net/http"
)

//Définissez une fonction main, qui sera la base de l'exécution de votre programme :
func main() {

	//Définissez trois HandleFunc (https://golang.org/pkg/net/http/#HandleFunc) et donnez-leur respectivement:
	//• Un path ”/”, et une fonction ”list”
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("/")
		return
	}

	http.HandleFunc("/", h1)

	//• Un path ”/done”, et une fonction ”done”
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("/done")
		return
	}

	http.HandleFunc("/done", h2)

	//• Un path ”/add”, et une fonction ”add”
	h3 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("/add")
		return
	}

	http.HandleFunc("/add", h3)

	//Appelez, dans le main, la fonction ListenAndServe (https://golang.org/pkg/net/http/#ListenAndServe) qui vous permettra de lancer votre serveur HTTP avec les handlers que vous venez de définir.
	log.Fatal(http.ListenAndServe(":8000", nil))
}

//Définissez une variable globale ”task” qui est une slice de Task.
var task []Task

//Définissez un type Task, qui est une struct avec deux champs:
type Task struct {

	//• ”Description”, de type string
	Description []string

	//• ”Done”, de type bool
	Done bool
}
