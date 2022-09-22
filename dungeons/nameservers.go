package dungeons

import (
	"fmt"
	"log"
	"net"

	"github.com/nanih98/dungeons/utils"
)

// GetDNSServers is a function that return the nameservers of the given domain
func GetDNSServers(domain string) []string {
	var nameservers []string
	nameserver, _ := net.LookupNS(domain)
	for _, ns := range nameserver {
		nameservers = append(nameservers, ns.Host)
	}
	return nameservers
}

// GetDNSIpv4 is a function that return the nameserver Ipv4 and Ipv6
func GetDNSIPS(nameserver string) (string, string) {
	ip, err := net.LookupIP(nameserver)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%s", ip[0]), fmt.Sprintf("%s", ip[1])
}

// DNSInfo prints the nameservers(+Ipv4) of the given domain
func DNSInfo(domain string) {
	log.Println("Checking nameservers for:", domain)
	var entry []string
	nameservers := GetDNSServers(domain)

	w := utils.TabWriter()
	fmt.Fprintln(w, "Nameserver\tIpv4\tIpv6")

	for _, nameserver := range nameservers {
		ipv4, ipv6 := GetDNSIPS(nameserver)
		entry = append(entry, nameserver, ipv4, ipv6)
		fmt.Fprintln(w, nameserver, "\t", ipv4, "\t", ipv6)
		entry = nil // flush the slice
	}
	w.Flush()
}
