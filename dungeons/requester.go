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
}

// Fetch the target
func Fetch(log *logger.CustomLogger, domain string, subdomains []string, ip string, workers int, logFormat string) {
	var errs []error

	workQueue := make(chan string, len(subdomains))

	wg := sync.WaitGroup{}
	wg.Add(workers)

	r := CustomResolver(ip)

	for i := 0; i < workers; i++ {
		worker := i
		go func(worker int, workQueue chan string) {
			for subdomain := range workQueue {
				start := time.Now()

				statusMsg, err := requester(domain, subdomain, r, ip)

				if err != nil {
					errs = append(errs, err)
				}

				request_time := fmt.Sprintf("%.2fs", time.Since(start).Seconds())
				outputLog(statusMsg, log, domain, ip, workers, logFormat, request_time)

				if statusMsg.Status == "Found" {
					outputLog(statusMsg, log, domain, ip, worker, logFormat, request_time)
				}

				//if statusMsg.Status == "Found" {
				//	log.CustomFields.Target = "lifullconnect.com"
				//	log.CustomFields.App = "dungeons"
				//	log.CustomFields.Domain = statusMsg.FullDomain
				//	log.CustomFields.Status = statusMsg.Status
				//	log.CustomFields.Seconds = fmt.Sprintf("%.2fs", time.Since(start).Seconds())
				//	log.CustomFields.Server = ip
				//	log.FuzzerFields()
				//	log.Info("Successfully scanned")
				//	//message := fmt.Sprintf("[Worker: %d] Server: %s ScannedDomain: %s | Status: %s | Time: (%.2fs)", worker, ip, statusMsg.FullDomain, statusMsg.Status, time.Since(start).Seconds())
				//	//log.Printf(message)
				//}
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

func requester(domain, subdomain string, r *net.Resolver, server string) (RecordInfo, error) {
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

func outputLog(statusMsg RecordInfo, log *logger.CustomLogger, domain string, ip string, worker int, logFormat string, time string) {
	if logFormat == "json" {
		log.CustomFields.Target = "lifullconnect.com"
		log.CustomFields.App = "dungeons"
		log.CustomFields.Domain = statusMsg.FullDomain
		log.CustomFields.Status = statusMsg.Status
		log.CustomFields.Seconds = time
		log.CustomFields.Server = ip
		log.FuzzerFields()
		log.Info("Successfully scanned")
		//message := fmt.Sprintf("[Worker: %d] Server: %s ScannedDomain: %s | Status: %s | Time: (%.2fs)", worker, ip, statusMsg.FullDomain, statusMsg.Status, time.Since(start).Seconds())
		//log.Printf(message)
	} else {
		message := fmt.Sprintf("[Worker: %d] Server: %s ScannedDomain: %s | Status: %s | Time: (%s)", worker, ip, statusMsg.FullDomain, statusMsg.Status, time)
		log.Info(message)
	}
}
