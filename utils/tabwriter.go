package utils

import (
	"os"
	"text/tabwriter"
)

func TabWriter() *tabwriter.Writer {
	w := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', tabwriter.Debug)
	return w
}
