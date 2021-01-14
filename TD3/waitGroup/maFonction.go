package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go maFonction(&wg)
	wg.Wait()
	fmt.Println("Fin du programme")
}

func maFonction(wg *sync.WaitGroup) {
	fmt.Println("J'ai fini")
	wg.Done()
}
