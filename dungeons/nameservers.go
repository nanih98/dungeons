package dungeons

import (
	"fmt"
	"log"
	"net"
)

func GetDNSServers(domain string) []string {
	var nameservers []string
	nameserver, _ := net.LookupNS(domain)
	for _, ns := range nameserver {
		nameservers = append(nameservers, ns.Host)
	}
	return nameservers
}

func Host(domain string) {
	hosts, err := net.LookupHost(domain)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hosts)
}
