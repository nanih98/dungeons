package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

func ReadFile() []string {
	var words []string
	//readFile, err := os.Open("/usr/local/share/SecLists/Discovery/Web-Content/directory-list-2.3-small.txt")
	readFile, err := os.Open("/usr/local/share/SecLists/Discovery/DNS/subdomains-top1million-5000.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}
	readFile.Close()

	return words
}

func CustomResolver(nameserver string) *net.Resolver {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, nameserver)
		},
	}
	return r
}

func main() {
	// Ejemplo de subdominios a consultar, pueden ser + de 10mil
	//var subdomains = []string{"status.edenor.com", "test.edenor.com", "www.edenor.com"}

	subdomains := ReadFile()

	// w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)

	// fmt.Fprintln(w, "Go type\tName\tTTL\tClass\tRR type\tetc")

	// fmt.Fprintf(w, "%T\t%[1]s\n", "[query, nameserver, status]")
	// w.Flush()

	resolver := CustomResolver("205.251.197.168:53")

	for _, subdomain := range subdomains {
		domain := subdomain + ".edenor.com"
		ip, err := resolver.LookupHost(context.Background(), domain)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Successfully found %s with the IP %s \n", domain, ip[0])
		}
	}
}

// var wg sync.WaitGroup

// // Requester block
// for i := 0; i < *maxPort; i++ {
// 	wg.Add(1)
// 	go func(i int) {
// 		Scanner(*ip, i)
// 		wg.Done()
// 	}(i)
// }

// wg.Wait()
// log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
