package cmd

import (
	"coffee-tracker/models"
	"coffee-tracker/storage"
	"fmt"

	"github.com/spf13/cobra"
)

// GetCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a coffee entry",
	Long: `Get coffee entries by GUID or name.
You can search for a specific coffee entry using either its GUID or name.
If searching by name, all matching entries will be returned.`,
	Run: func(cmd *cobra.Command, args []string) {
		guid, _ := cmd.Flags().GetString("guid")
		name, _ := cmd.Flags().GetString("name")

		store := storage.NewStore()

		if guid != "" {
			coffee, err := store.GetByGUID(guid)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			printCoffee(coffee)
		} else if name != "" {
			coffees, err := store.GetByName(name)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			for _, coffee := range coffees {
				printCoffee(coffee)
			}
		} else {
			fmt.Println("Error: specify --guid or --name")
			cmd.Help()
		}
	},
}

func init() {
	// Add flags
	GetCmd.Flags().StringP("guid", "g", "", "GUID of the coffee entry")
	GetCmd.Flags().StringP("name", "n", "", "Name of the coffee to search")

	// Add validation to ensure at least one search parameter is provided
	GetCmd.MarkFlagsOneRequired("guid", "name")
}

func printCoffee(coffee *models.Coffee) {
	fmt.Printf("\nCoffee: %s\n", coffee.Name)
	fmt.Printf("GUID: %s\n", coffee.GUID)
	fmt.Printf("Notes: %s\n", coffee.Notes)
	fmt.Printf("Duration: %ds | In: %dg | Out: %dg\n",
		coffee.Duration, coffee.In, coffee.Out)
	fmt.Printf("Date: %s\n", coffee.Date.Format("2006-01-02 15:04"))
}
