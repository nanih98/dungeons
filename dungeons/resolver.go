package dungeons

import (
	"context"
	"net"
	"time"
)

func ResolverIPS(domain string) []string {
	var ips []string

	nameservers := GetDNSServers(domain)

	for _, nameserver := range nameservers {
		ipv4, _ := GetDNSIPS(nameserver)
		ips = append(ips, ipv4)
	}
	return ips
}

// CustomResolver
func CustomResolver(server string) *net.Resolver {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, server+":53")
		},
	}

	// ip, err := r.LookupHost(context.Background(), "ñiñiñiñ.google.com")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	return r
}
