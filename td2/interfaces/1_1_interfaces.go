package main

import "fmt"

type IPAddr [4]byte

func main() {
	// Use IPAddr interface
	var ip = IPAddr{1, 2, 3, 4}
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
func (ipAd *IPAddr)String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAd[0], ipAd[1], ipAd[2], ipAd[3])
}
