package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func main1() {
	// Use IPAddr interface
	var ip = IPAddr{1, 2, 3, 4}
	// fmt.Println(ip.String())
	fmt.Println(ip.String())

	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	// for name, ip := range hosts {
	// 	fmt.Printf("%v: %v\n", name, ip)
	// }
}

// Implement fmt.Stringer interface
// func (ipAd *IPAddr)String() string {
// 	return fmt.Sprintf("%v.%v.%v.%v", ipAd[0], ipAd[1], ipAd[2], ipAd[3])
// }

// Implement interface via strconv
func (ipAd *IPAddr)String() string {
	return convert(ipAd) 
}

// strconv
func convert( b *IPAddr ) string { // IPAddr sans étoile
	s := make([]string,len(b))
	for i := range b {
			s[i] = strconv.Itoa(int(b[i]))
	}
	// for i, val := range b {
	// 		s[i] = strconv.Itoa(int(val))
	// }
	// for i := 0; i < len(b); i++ {
	// 	s[i] = strconv.Itoa(int(b[i]))
	// }
	return strings.Join(s,".")
}