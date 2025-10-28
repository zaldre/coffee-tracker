package cmd

import (
	"coffee-tracker/storage"
	"fmt"

	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List coffee entries",
	Long: `List your coffee brewing sessions.
This command shows a summary of your coffee entries with basic information
including name, GUID, and date. Use the --limit flag to control how many
entries are displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")

		store := storage.NewStore()
		coffees, err := store.List(limit)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("\nShowing %d coffee entries:\n", len(coffees))
		for _, coffee := range coffees {
			fmt.Printf("- %s (%s) - %s\n",
				coffee.Name,
				coffee.GUID[:8],
				coffee.Date.Format("2006-01-02"))
		}
	},
}

func init() {
	// Add flags
	ListCmd.Flags().IntP("limit", "l", 10, "Number of entries to show")
}
