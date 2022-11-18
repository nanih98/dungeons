package dungeons

import (
	"fmt"
	"github.com/nanih98/dungeons/logger"
	"net"
)

// GetIPV4 return the ipv4 of the given domain (in string format ex: ns1.example.com)
func GetIPV4(server string, log *logger.CustomLogger) string {
	ip, err := net.LookupIP(server)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", ip[0])
}

// GetIPV6 return the ipv6 of the given domain (in string format ex: ns1.example.com)
func GetIPV6(server string, log *logger.CustomLogger) string {
	ip, err := net.LookupIP(server)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%v", ip[1])
}

// GetNameservers returns the list of all the nameservers of the given domain
func GetNameservers(domain string) []string {
	var nameservers []string
	nameserver, _ := net.LookupNS(domain)
	for _, ns := range nameserver {
		nameservers = append(nameservers, ns.Host)
	}
	return nameservers
}
