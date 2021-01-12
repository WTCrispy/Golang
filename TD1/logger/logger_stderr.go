package main

import (
	"log"
	"os"
)

func main(){
	logger := log.New(os.Stderr,"Custom logger:",log.LstdFlags)
	logger.Print("toto")
}
