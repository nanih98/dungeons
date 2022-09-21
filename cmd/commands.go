package cmd

import (
	"runtime"

	"github.com/nanih98/dungeons/dungeons"
	"github.com/spf13/cobra"
)

var (
	version   string
	goversion = runtime.Version()
	goos      = runtime.GOOS
	goarch    = runtime.GOARCH
)

func Info(domain *string) *cobra.Command {
	return &cobra.Command{
		Use:   "info",
		Short: "Info of the nameservers to the given domain",
		Long:  "Check the nameservers of the given domain in the cli",
		Run: func(cmd *cobra.Command, args []string) {
			dungeons.DNSInfo(*domain)
		},
	}
}

func Fuzz(domain *string) *cobra.Command {
	return &cobra.Command{
		Use:   "fuzz",
		Short: "Start massive requests to all the nameservers.",
		Long:  "Start massive requests to all the nameservers of the given domain using a dictionary",
		Run: func(cmd *cobra.Command, args []string) {
			dungeons.DNSInfo(*domain)
		},
	}
}
