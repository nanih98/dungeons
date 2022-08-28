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

func Nameservers(domain *string) *cobra.Command {
	return &cobra.Command{
		Use:   "nameservers",
		Short: "Check nameservers to the given domain",
		Long:  "Check the nameservers of the given domain in the cli",
		Run: func(cmd *cobra.Command, args []string) {
			dungeons.DNSInfo(*domain)
		},
	}
}
