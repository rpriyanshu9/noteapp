package cmd

import (
	"fmt"

	"github.com/rpriyanshu9/noteapp/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all notes",
	Long:  "Lists all notes of the user",
	Run: func(cmd *cobra.Command, args []string) {

		// Loading existing notes from the note.json file
		res := utils.LoadNotes()

		// Checking if no notes present
		if len(res) == 0 {
			fmt.Println("No notes to display.")
			return
		}

		// Printing notes
		fmt.Printf("\nYour notes are: \n\n")
		for _, m := range res {
			fmt.Printf("-> %s: %s\n", m.Title, m.Body)
		}
		fmt.Printf("\n")
	},
}
