package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	updateID    int
	updateNom   string
	updateEmail string
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing contact",
	Long:  `Update an existing contact's name or email by providing its ID.`, 
	Run: func(cmd *cobra.Command, args []string) {
		if updateID == 0 {
			fmt.Println("The --id flag is required to update a contact.")
			return
		}

		contact, err := dataStore.Get(uint(updateID))
		if err != nil {
			fmt.Println("Contact not found:", err)
			return
		}

		if updateNom != "" {
			contact.Nom = strings.TrimSpace(updateNom)
		}

		if updateEmail != "" {
			contact.Email = strings.TrimSpace(updateEmail)
		}

		if err := dataStore.Update(contact); err != nil {
			fmt.Println("Error updating contact:", err)
		} else {
			fmt.Println("Contact updated successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID of the contact to update")
	updateCmd.Flags().StringVarP(&updateNom, "nom", "n", "", "New name of the contact (leave empty to keep current)")
	updateCmd.Flags().StringVarP(&updateEmail, "email", "e", "", "New email of the contact (leave empty to keep current)")
}
