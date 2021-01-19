package main

import (
	"fmt"
	//	"strconv"
	//	"strings"
	"time"
)

type IPAddr struct {
	When time.Time
	What string
}

//func test() {
//	hosts := map[string]IPAddr{
//		"loopback":  {127, 0, 0, 1},
//		"googleDNS": {8, 8, 8, 8},
//	}
//
//	for _, ip := range hosts {
//		res := []string{}
//		for _, val := range ip {
//			res = append(res, strconv.Itoa(int(val)))
//		}
//		fmt.Printf("%v\n", strings.Join(res, "."))
//	}
//}

func (err *IPAddr) Error() string {
	return fmt.Sprintf("When %v: \nWhat: %v", err.When, err.What)
}

func run() error {
	return &IPAddr{
		When: time.Now(),
		What: "Toto",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		return
	}

}
