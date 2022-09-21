/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dungeons",
	Short: "Massive DNS requests",
	Long:  `Massive DNS requests. DNS fuzzing`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute starts the cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("😢 %s\n", err.Error())
		os.Exit(1)
	}
}
