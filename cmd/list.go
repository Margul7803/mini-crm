package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contacts",
	Long:  `List all contacts currently stored in the CRM.`,
	Run: func(cmd *cobra.Command, args []string) {
		contacts, err := dataStore.GetAll()
		if err != nil {
			fmt.Println("Error retrieving contacts:", err)
			return
		}

		fmt.Println("--- List of Contacts ---")
		if len(contacts) == 0 {
			fmt.Println("No contacts found.")
			return
		}
		for _, c := range contacts {
			fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

