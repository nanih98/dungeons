package cmd

import (
	"github.com/nanih98/dungeons/logger"
)

var (
	domain     string
	workers    int
	output     string
	dictionary string
	level      string
	logFormat  string
)

func init() {
	log := logger.Logger()

	info := Info(&domain, &output, &log, &level)
	fuzz := Fuzz(&domain, &workers, &dictionary, &log, &level, &logFormat)

	rootCmd.AddCommand(info)
	rootCmd.AddCommand(fuzz)

	// Info
	info.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	info.PersistentFlags().StringVar(&output, "output", "tabwriter", "Output mode. Tabwriter or json. Default: tabwriter")
	info.PersistentFlags().StringVar(&level, "level", "info", "Setup log level")

	// Fuzzer
	fuzz.PersistentFlags().StringVar(&domain, "domain", "", "Enter the domain")
	fuzz.PersistentFlags().IntVar(&workers, "workers", 5, "Enter the max workers (threads)")
	fuzz.PersistentFlags().StringVar(&dictionary, "dictionary", "", "Dictionary path")
	fuzz.PersistentFlags().StringVar(&logFormat, "logFormat", "text", "Log format: Text or Json")
	fuzz.PersistentFlags().StringVar(&level, "level", "info", "Setup log level")

	// Required flags
	info.MarkPersistentFlagRequired("domain")
	fuzz.MarkPersistentFlagRequired("domain")
	fuzz.MarkPersistentFlagRequired("dictionary")
}
