package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	strs := make([]string, len(ip))

	for i := 0; i < len(ip); i++ {
		strs[i] = strconv.Itoa(int(ip[i]))
	}

	return strings.Join(strs, ".")
}

// func (ip IPAddr) String() string {
//     return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
// }

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
