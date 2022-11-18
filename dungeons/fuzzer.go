package dungeons

import (
	"context"
	"fmt"
	"github.com/nanih98/dungeons/logger"
	"net"
	"sync"
	"time"
)

type RecordInfo struct {
	FullDomain string
	Status     string
	Ips        []string
	Time       string
	Error      error
}

// Fetch all the subdomains to the given domain
func Fetch(log *logger.CustomLogger, domain string, subdomains []string, ip string, workers int, logFormat string) {
	var errs []error

	workQueue := make(chan string, len(subdomains))

	wg := sync.WaitGroup{}
	wg.Add(workers)

	r := customResolver(ip)

	for i := 0; i < workers; i++ {
		worker := i
		go func(worker int, workQueue chan string) {
			for subdomain := range workQueue {
				result := requester(domain, subdomain, r, ip)

				if result.Error != nil {
					errs = append(errs, result.Error)
				}

				// Only show valid results
				if result.Status == "Found" {
					outputLog(result, log, ip, workers, logFormat, result.Time)
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

func requester(domain, subdomain string, r *net.Resolver, server string) RecordInfo {
	start := time.Now()

	fullDomain := subdomain + "." + domain

	// Fetch if the subdomain exists
	ipHost, err := r.LookupHost(context.Background(), fullDomain)

	requestTime := fmt.Sprintf("%.2fs", time.Since(start).Seconds())

	if err != nil {
		return RecordInfo{
			FullDomain: fullDomain,
			Status:     "NotFound",
			Ips:        ipHost,
			Time:       requestTime,
			Error:      err,
		}
	}

	return RecordInfo{
		FullDomain: fullDomain,
		Status:     "Found",
		Ips:        ipHost,
		Time:       requestTime,
		Error:      nil,
	}
}

func outputLog(recordInfo RecordInfo, log *logger.CustomLogger, nameserver string, worker int, logFormat string, time string) {
	if logFormat == "json" {
		log.CustomFields.FullDomain = recordInfo.FullDomain
		log.CustomFields.Status = recordInfo.Status
		log.CustomFields.Ips = recordInfo.Ips
		log.CustomFields.RecordTime = time
		log.CustomFields.Nameserver = nameserver
		log.FuzzerFields()
		log.Info(recordInfo.Status)
	} else {
		message := fmt.Sprintf("[Worker: %d] Server: %s ScannedDomain: %s | Status: %s | Time: (%s)", worker, nameserver, recordInfo.FullDomain, recordInfo.Status, time)
		log.Info(message)
	}
}

func customResolver(server string) *net.Resolver {
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
