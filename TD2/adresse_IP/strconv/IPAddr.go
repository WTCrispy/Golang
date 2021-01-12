package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for _, ip := range hosts {
		res := []string{}
		for _, val := range ip {
			res = append(res, strconv.Itoa(int(val)))
		}
		fmt.Printf("%v\n", strings.Join(res, "."))
	}
}
