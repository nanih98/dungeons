package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"text/tabwriter"
)

type nameservers []string

func GetNameservers(domain string) nameservers {
	nameservers := nameservers{}

	nameserver, _ := net.LookupNS(domain)
	for _, ns := range nameserver {
		nameservers = append(nameservers, ns.Host)
	}
	return nameservers
}

func (n *nameservers) GetIPV4() []string {
	var ipv4 []string
	for _, nameserver := range *n {
		ip, err := net.LookupIP(nameserver)
		if err != nil {
			log.Println(err)
		}
		ipv4 = append(ipv4, fmt.Sprintf("%v", ip[0]))
	}
	return ipv4
}

func (n *nameservers) GetIPV6() []string {
	var ipv6 []string
	for _, nameserver := range *n {
		ip, err := net.LookupIP(nameserver)
		if err != nil {
			log.Println(err)
		}

		ipv6 = append(ipv6, fmt.Sprintf("%v", ip[1]))
	}

	return ipv6
}

func TabWriter() *tabwriter.Writer {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', tabwriter.Debug)
	return w
}

func (n *nameservers) Info() {
	fmt.Println(*n) // prints list of nameservers

	// log.Printf("Checking nameservers for: %s \n\n", domain)
	// var entry []string
	// w := utils.TabWriter()
	// fmt.Fprintln(w, "Nameserver\t Ipv4\t Ipv6")

	// for _, nameserver := range *n {
	// 	ipv4 := n.GetIPV4()
	// 	ipv6 := n.GetIPV6()
	// 	entry = append(entry, nameserver, ipv4, ipv6)
	// 	fmt.Fprintln(w, nameserver, "\t", ipv4, "\t", ipv6)
	// 	entry = nil
	// }
	// w.Flush()
}

func main() {
	nameservers := GetNameservers("edrans.com")
	nameservers.Info()

}
