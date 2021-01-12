package main

import (
	"fmt"
	"flag"
	"time"
	"strings"
	"net/http"
	"log"
	"strconv"
	"io/ioutil"
)

func time() {
    start := time.Now()

    r := new(big.Int)
    fmt.Println(r.Binomial(1000, 10))

    elapsed := time.Since(start)
    log.Printf("Binomial took %s", elapsed)
}

func main(){
	url := flag.String("url", "", "list of urls")
	flag.Parse()
	if *url != "" {
		list_url := strings.Split(*url,",")
		for index, element := range list_url {
			fmt.Println(element)
			resp, err := http.Get(element)
			if err != nil {
				log.Fatal(err)
				continue
			}
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil{
				log.Fatal(err)
			}
			defer resp.Body.Close()
			filename := "image-"+strconv.Itoa(index)+".jpeg"
			ioutil.WriteFile(filename, data, 0644)
		}
		return
	}
	fmt.Print("urls is empty")
}
