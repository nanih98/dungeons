package cmd

var (
	domain  string
	workers int
	output  string
)

func init() {
	info := Info(&domain, &output)
	fuzz := Fuzz(&domain)

	rootCmd.AddCommand(info)
	rootCmd.AddCommand(fuzz)

	// Info
	info.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	info.PersistentFlags().StringVar(&output, "output", "tabwriter", "Output mode. Tabwriter or json. Default: tabwriter")

	// Fuzzer
	fuzz.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	fuzz.PersistentFlags().IntVar(&workers, "workers", 10, "Enter the max workers (threads)")

	// Required flags
	info.MarkPersistentFlagRequired("domain")
	fuzz.MarkPersistentFlagRequired("domain")
}
