package dungeons

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type RecordInfo struct {
	FullDomain string
	Status     string
	Ips        []string
}

// Start the fuzzer
func Fetch(domain string, subdomains []string, ip string, workers int) {
	var errs []error

	workQueue := make(chan string, len(subdomains))

	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		worker := i
		go func(worker int, workQueue chan string) {
			for subdomain := range workQueue {
				start := time.Now()

				statusMsg, err := requester(domain, subdomain, ip)

				if err != nil {
					errs = append(errs, err)
				}

				if statusMsg.Status == "Found" {
					message := fmt.Sprintf("[Worker: %d] Server: %s ScannedDomain: %s | Status: %s | Time: (%.2fs)", worker, ip, statusMsg.FullDomain, statusMsg.Status, time.Since(start).Seconds())
					log.Printf(message)
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

	return r
}

func requester(domain, subdomain, server string) (RecordInfo, error) {
	r := CustomResolver(server)

	fullDomain := subdomain + "." + domain

	ip, err := r.LookupHost(context.Background(), fullDomain)

	if err != nil {
		return RecordInfo{
			FullDomain: fullDomain,
			Status:     "NotFound",
			Ips:        ip,
		}, err
	}

	statusMsg := RecordInfo{
		FullDomain: fullDomain,
		Status:     "Found",
		Ips:        ip,
	}

	return statusMsg, nil
}
