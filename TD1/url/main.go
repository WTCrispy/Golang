package main

import (
	"fmt"
	"flag"
	"strings"
	"net/http"
	"log"
	"strconv"
	"io/ioutil"
	//"os"
)

func main(){
	// default_url = "http:url1.tld,http://url2.tld"
	urls := flag.String("urls", "", "list of urls")
	flag.Parse()
	if *urls != "" {
		list_url := strings.Split(*urls,",")
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
