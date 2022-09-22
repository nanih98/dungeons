package dungeons

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Fetch function
func Fetch(subdomains, ips []string, workers int) {
	var errs []error

	workQueue := make(chan string, len(subdomains))

	wg := sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		worker := i
		go func(worker int, workQueue chan string) {
			for subdomain := range workQueue {
				start := time.Now()

				status, err := requester(subdomain, ips[0])

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

func requester(subdomain, server string) (string, error) {
	r := CustomResolver(server)

	fullDomain := subdomain + ".vpnroulette.com"

	ip, err := r.LookupHost(context.Background(), fullDomain)

	if err != nil {
		//return fmt.Sprintf("Host not found %s", fullDomain), nil
		return "", nil
	}
	return fmt.Sprintf("Host found %s with the IP(s) %v", fullDomain, ip), nil
}
