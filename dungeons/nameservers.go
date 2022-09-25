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

// GetDNSIPS is a function that return the nameserver Ipv4 and Ipv6
func GetDNSIPS(nameserver string) (string, string) {
	ip, err := net.LookupIP(nameserver)
	if err != nil {
		log.Println(err)
	}

	if len(ip) > 1 {
		return fmt.Sprintf("%s", ip[0]), fmt.Sprintf("%s", ip[1])
	}
	return fmt.Sprintf("%s", ip[0]), "Null"
}

// ResolverIPV4
func ResolverIPV4(domain string) []string {
	var ips []string

	nameservers := GetDNSServers(domain)

	for _, nameserver := range nameservers {
		ipv4, _ := GetDNSIPS(nameserver)
		ips = append(ips, ipv4)
	}
	return ips
}

// Info prints information about the nameservers
func Info(domain string) {
	log.Printf("Checking nameservers for: %s \n\n", domain)
	var entry []string
	nameservers := GetDNSServers(domain)

	w := utils.TabWriter()
	fmt.Fprintln(w, "Nameserver\t Ipv4\t Ipv6")

	for _, nameserver := range nameservers {
		ipv4, ipv6 := GetDNSIPS(nameserver)
		entry = append(entry, nameserver, ipv4, ipv6)
		fmt.Fprintln(w, nameserver, "\t", ipv4, "\t", ipv6)
		entry = nil
	}
	w.Flush()
}
