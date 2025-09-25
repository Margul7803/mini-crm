package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	deleteID int
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a contact",
	Long:  `Delete a contact from the CRM by providing its ID.`, 
	Run: func(cmd *cobra.Command, args []string) {
		if deleteID == 0 {
			fmt.Println("The --id flag is required to delete a contact.")
			return
		}

		if err := dataStore.Delete(uint(deleteID)); err != nil {
			fmt.Println("Error deleting contact:", err)
		} else {
			fmt.Println("Contact deleted successfully!")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID of the contact to delete")
}
