package cmd

var (
	domain  string
	workers int
	output  string
	path    string
)

func init() {
	info := Info(&domain, &output)
	fuzz := Fuzz(&domain, &workers, &path)

	rootCmd.AddCommand(info)
	rootCmd.AddCommand(fuzz)

	// Info
	info.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	info.PersistentFlags().StringVar(&output, "output", "tabwriter", "Output mode. Tabwriter or json. Default: tabwriter")

	// Fuzzer
	fuzz.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	fuzz.PersistentFlags().IntVar(&workers, "workers", 5, "Enter the max workers (threads)")
	fuzz.PersistentFlags().StringVar(&path, "path", "", "Dictionary path")

	// Required flags
	info.MarkPersistentFlagRequired("domain")
	fuzz.MarkPersistentFlagRequired("domain")
	fuzz.MarkPersistentFlagRequired("path")
}
