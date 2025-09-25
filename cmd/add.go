package cmd

import (
	"fmt"
	"strings"

	"mini-crm/store"

	"github.com/spf13/cobra"
)

var (
	addNom   string
	addEmail string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new contact",
	Long:  `Add a new contact to the CRM with a name and email.`, 
	Run: func(cmd *cobra.Command, args []string) {
		if addNom == "" || addEmail == "" {
			fmt.Println("Both --nom and --email flags are required.")
			return
		}

		contact := store.Contact{
			Nom:   strings.TrimSpace(addNom),
			Email: strings.TrimSpace(addEmail),
		}

		if err := dataStore.Add(contact); err != nil {
			fmt.Println("Error adding contact:", err)
		} else {
			fmt.Println("Contact added successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addNom, "nom", "n", "", "Name of the contact")
	addCmd.Flags().StringVarP(&addEmail, "email", "e", "", "Email of the contact")
}
