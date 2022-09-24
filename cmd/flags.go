package cmd

var (
	domain  string
	workers int
)

func init() {
	info := Info(&domain)
	fuzz := Fuzz(&domain)

	rootCmd.AddCommand(info)
	rootCmd.AddCommand(fuzz)

	// Info
	info.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")

	// Fuzzer
	fuzz.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	fuzz.PersistentFlags().IntVar(&workers, "workers", 10, "Enter the max workers (threads)")

	// Required flags
	info.MarkPersistentFlagRequired("domain")
	fuzz.MarkPersistentFlagRequired("domain")
}
