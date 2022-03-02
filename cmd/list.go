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
		res := utils.LoadNotes()
		if len(res) == 0 {
			fmt.Println("No notes to display.")
			return
		}
		fmt.Printf("\nYour notes are: \n\n")
		for _, m := range res {
			fmt.Printf("-> %s: %s\n", m["title"], m["body"])
		}
		fmt.Printf("\n")
	},
}
