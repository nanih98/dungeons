package main

import (
	"context"
	"net"
	"time"
)

func main() {
	// Custom resolver
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "1.1.1.1:53")
		},
	}

	ip, _ := r.LookupHost(context.Background(), "www.google.com")

	print(ip[0])
}
