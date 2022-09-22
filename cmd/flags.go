package cmd

var (
	profileName string
	startURL    string
	region      string
	level       string
)

var (
	domain string
)

func init() {
	info := Info(&domain)
	fuzz := Fuzz(&domain)

	rootCmd.AddCommand(info)
	rootCmd.AddCommand(fuzz)

	// Nameservers
	info.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	info.MarkPersistentFlagRequired("domain")

	fuzz.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	fuzz.MarkPersistentFlagRequired("domain")
}
