package cmd

import (
	"fmt"

	"github.com/rpriyanshu9/noteapp/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read title",
	Short: "Reads a note",
	Long:  "\nReads a note for the title given by user",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Checking for any errors. Specifically the "too few/too many arguments" error
		if err := readErrFunction(args); err != nil {
			return err
		}

		// Getting title from arguments provided
		title := args[0]

		// Loading existing notes from the note.json file
		notes := utils.LoadNotes()

		// Searching for the note which matches the title
		for _, m := range notes {
			if m["title"] == title {
				fmt.Printf("\nYour note is: \n-> %s: %s\n\n", m["title"], m["body"])
				return nil
			}
		}

		// If we reach here, note with the provided title doesn't exist
		fmt.Printf("\nError: no note with title: \"%s\" present\n\n", title)
		return nil
	},
}

// Checking if the number of arguments provided are as expected
func readErrFunction(args []string) error {
	if len(args) < 1 || len(args) > 1 {
		return fmt.Errorf("1 argument expected")
	}
	return nil
}
