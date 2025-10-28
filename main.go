package main

import (
	"coffee-tracker/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const VERSION = "0.0.1"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "coffee-tracker",
		Short: "A coffee brewing tracker application",
		Long: `Coffee Tracker is a CLI application for tracking your coffee brewing sessions.
It helps you record details about your coffee preparation including grind settings,
timing, ratios, and tasting notes.`,
		Version: VERSION,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	// Add subcommands
	rootCmd.AddCommand(cmd.AddCmd)
	rootCmd.AddCommand(cmd.GetCmd)
	rootCmd.AddCommand(cmd.ListCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
