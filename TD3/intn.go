package main

import (
	//"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(time.Now().UnixNano())
  //      fmt.Println(rand.Intn(100))
    //    fmt.Println(rand.Intn(100))
      //  fmt.Println(rand.Intn(100))
ch:= make(chan interface{})
func(){
rand.Intn(100)
}()
}
