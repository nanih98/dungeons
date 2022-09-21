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

	rootCmd.AddCommand(info)

	// Nameservers
	info.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	info.MarkPersistentFlagRequired("domain")
}
