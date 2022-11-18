package cmd

import (
	"github.com/nanih98/dungeons/dungeons"
	"github.com/nanih98/dungeons/logger"
	"github.com/nanih98/dungeons/utils"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	version   string
	goversion = runtime.Version()
	goos      = runtime.GOOS
	goarch    = runtime.GOARCH
)

//func Info(domain *string, output *string, log *logger.CustomLogger, level *string) *cobra.Command {
//	return &cobra.Command{
//		Use:   "info",
//		Short: "Info of the nameservers to the given domain",
//		Long:  "Check the nameservers of the given domain in the cli",
//		Run: func(cmd *cobra.Command, args []string) {
//			log.LogLevel(*level)
//			dungeons.Info(log, *domain, *output)
//		},
//	}
//}

func Fuzz(domain *string, workers *int, dictionary *string, log *logger.CustomLogger, level *string, logFormat *string) *cobra.Command {
	return &cobra.Command{
		Use:   "fuzz",
		Short: "Start fuzzing DNS servers",
		Long:  "Start the fuzzer using the subdomain words of the given dictionary.",
		Run: func(cmd *cobra.Command, args []string) {
			log.LogLevel(*level)
			log.LogFormat(*logFormat)
			nameservers := dungeons.GetNameservers(*domain)
			subdomains := utils.ReadDictionary(log, *dictionary)

			// Start querying
			// Get a random NS server of the domain
			ip := dungeons.GetIPV4(utils.RandomServer(nameservers), log)
			dungeons.Fetch(log, *domain, subdomains, ip, *workers, *logFormat)
		},
	}
}

// MassiveRequests is a function that execute massive requests to all the nameservers of the given domain
//func MassiveRequests(domain *string, workers *int, dictionary *string, log *logger.CustomLogger, level *string, logFormat *string) *cobra.Command {
//	return &cobra.Command{
//		Use:   "massive",
//		Short: "Start massive requests to all the nameservers.",
//		Long:  "Start massive requests to all the nameservers of the given domain using a dictionary",
//		Run: func(cmd *cobra.Command, args []string) {
//			log.LogLevel(*level)
//			log.LogFormat(*logFormat)
//			nameservers := dungeons.GetNameservers(*domain)
//			subdomains := utils.ReadDictionary(log, *dictionary)
//			log.Info(fmt.Sprintf("Using %d workers", *workers))
//
//			var wg sync.WaitGroup
//
//			for _, server := range nameservers {
//				ip := dungeons.GetIPV4(server, log)
//				wg.Add(1)
//				go func(subdomains []string, ip string) {
//					defer wg.Done()
//					dungeons.Fetch(log, *domain, subdomains, ip, *workers, *logFormat)
//				}(subdomains, ip)
//			}
//			wg.Wait()
//		},
//	}
//}

//var wg sync.WaitGroup
//
//for _, server := range nameservers {
//	ip := dungeons.GetIPV4(server, log)
//	wg.Add(1)
//	go func(subdomains []string, ip string) {
//		defer wg.Done()
//		dungeons.Fetch(log, *domain, subdomains, ip, *workers, *logFormat)
//	}(subdomains, ip)
//}
//wg.Wait()
