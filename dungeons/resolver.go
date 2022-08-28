package dungeons

import (
	"context"
	"fmt"
	"net"
	"time"
)

// CustomResolver is a function that...
func CustomResolver() {
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

	ip, err := r.LookupHost(context.Background(), "ñiñiñiñ.google.com")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ip)
}
