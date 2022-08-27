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
	nameservers := Nameservers(&domain)

	rootCmd.AddCommand(nameservers)

	// Nameservers
	nameservers.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	nameservers.MarkPersistentFlagRequired("domain")
}
