package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	subdomains := readFile()

	log.Println(len(subdomains), "subdomains to check")
	//resolver := customResolver("205.251.197.168:53")

	var w int

	flag.IntVar(&w, "w", 1, "amount of workers")
	flag.Parse()

	fetch(subdomains, w)

}

func fetch(subdomains []string, workers int) {
	var errs []error

	workQueue := make(chan string, len(subdomains))

	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		worker := i
		go func(worker int, workQueue chan string) {
			for subdomain := range workQueue {
				start := time.Now()

				status, err := resolver(subdomain)

				if err != nil {
					errs = append(errs, err)
				}

				if status != "" {
					log.Printf("[Worker: %d] Status: %s (%.2fs)",
						worker,
						status,
						time.Since(start).Seconds())
				}
			}
			wg.Done()
		}(worker, workQueue)
	}

	go func() {
		for _, u := range subdomains {
			workQueue <- u
		}
		close(workQueue)
	}()
	wg.Wait()
}

func resolver(subdomain string) (string, error) {
	r := customResolver("8.8.8.8:53")

	// r := &net.Resolver{
	// 	PreferGo: true,
	// 	Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
	// 		d := net.Dialer{
	// 			Timeout: time.Millisecond * time.Duration(10000),
	// 		}
	// 		return d.DialContext(ctx, network, "8.8.8.8:53")
	// 	},
	// }

	fullDomain := subdomain + ".edenor.com"

	ip, err := r.LookupHost(context.Background(), fullDomain)

	if err != nil {
		//return fmt.Sprintf("Host not found %s", fullDomain), nil
		return "", nil
	}
	return fmt.Sprintf("Host found %s with the IP(s) %v", fullDomain, ip), nil
}

func customResolver(nameserver string) *net.Resolver {
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

func readFile() []string {
	log.Println("Reading subdomains from /usr/local/share/SecLists/Discovery/DNS/subdomains-top1million-20000.txt")
	var words []string
	readFile, err := os.Open("/usr/local/share/SecLists/Discovery/DNS/subdomains-top1million-20000.txt")

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
