package cmd

import (
	"log"
	"runtime"
	"sync"

	"github.com/nanih98/dungeons/dungeons"
	"github.com/nanih98/dungeons/utils"
	"github.com/spf13/cobra"
)

var (
	version   string
	goversion = runtime.Version()
	goos      = runtime.GOOS
	goarch    = runtime.GOARCH
)

func Info(domain *string, output *string) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Info of the nameservers to the given domain",
		Long:  "Check the nameservers of the given domain in the cli",
		Run: func(cmd *cobra.Command, args []string) {
			dungeons.Info(*domain, *output)
		},
	}
}

func Fuzz(domain *string, workers *int, path *string) *cobra.Command {
	return &cobra.Command{
		Use:   "fuzz",
		Short: "Start massive requests to all the nameservers.",
		Long:  "Start massive requests to all the nameservers of the given domain using a dictionary",
		Run: func(cmd *cobra.Command, args []string) {
			nameservers := dungeons.GetNameservers(*domain)
			subdomains := utils.ReadDictionary(*path)
			log.Printf("Using %d workers", *workers)
			var wg sync.WaitGroup

			for _, server := range nameservers {
				ip := dungeons.GetIPV4(server)
				wg.Add(1)
				go func(subdomains []string, ip string) {
					defer wg.Done()
					dungeons.Fetch(*domain, subdomains, ip, *workers)
				}(subdomains, ip)
			}
			wg.Wait()
		},
	}
}
