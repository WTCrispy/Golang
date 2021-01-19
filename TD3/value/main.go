package main

import (
	"fmt"
	//"sync"
	"math/rand"
	"time"

)
type myError struct {
	What string
	Value int
}
	
func (me *myError) Error() string {
    return fmt.Sprintf("Erreur: %v (valeur: %v)", me.What, me.Value)
}

func main(){
	c := make(chan interface{})
	go func (){
		rand.Seed(time.Now().UnixNano())
		value := rand.Intn(100)
		if value < 50 {
			c <- &myError{"Valeur inférieur à 50",value}
			return
		}
		c <- value
	}()
	result := <-c
	switch v := result.(type) {
	case int:
		fmt.Printf("Le nombre suivant est bon: %d\n", v)
	case error:
		fmt.Println(v)
	}
	
}
