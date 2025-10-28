package cmd

import (
	"coffee-tracker/models"
	"coffee-tracker/storage"
	"fmt"

	"github.com/spf13/cobra"
)

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new coffee entry",
	Long: `Add a new coffee brewing session to your tracker.
This command allows you to record details about your coffee preparation including
grind settings, timing, ratios, and tasting notes.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		notes, _ := cmd.Flags().GetString("notes")
		duration, _ := cmd.Flags().GetInt("duration")
		in, _ := cmd.Flags().GetInt("in")
		out, _ := cmd.Flags().GetInt("out")
		grind, _ := cmd.Flags().GetInt("grind")
		rating, _ := cmd.Flags().GetInt("rating")

		if name == "" {
			fmt.Println("Error: --name is required")
			cmd.Help()
			return
		}

		coffee := models.NewCoffee(name, notes, duration, in, out, grind, rating)

		store := storage.NewStore()
		if err := store.Save(coffee); err != nil {
			fmt.Printf("Error saving coffee: %v\n", err)
			return
		}

		fmt.Printf("Added coffee: %s (GUID: %s)\n", coffee.Name, coffee.GUID)
	},
}

func init() {
	// Add flags
	AddCmd.Flags().StringP("name", "n", "", "The name of the coffee (required)")
	AddCmd.Flags().String("notes", "", "Notes about the coffee")
	AddCmd.Flags().IntP("duration", "d", 0, "Duration of shot in seconds")
	AddCmd.Flags().IntP("in", "i", 0, "Weight in grams of ground coffee")
	AddCmd.Flags().IntP("out", "o", 0, "Weight of coffee output/yield in grams")
	AddCmd.Flags().IntP("grind", "g", 0, "Grind setting on the grinder")
	AddCmd.Flags().IntP("rating", "r", 0, "Rating from 1-5 of the brew")

	// Mark required flags
	AddCmd.MarkFlagRequired("name")
}
