package cmd

import (
	"fmt"

	"github.com/rpriyanshu9/noteapp/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add title body",
	Short: "Adds a new note or Updates existing note",
	Long:  "Adds a new note with title and body given by the user or update already present note",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Checking for any errors. Specifically the "too few/too many arguments" error
		if err := addErrFunc(args); err != nil {
			return err
		}

		// Getting title and body of the note from the arguments provided
		title, body := args[0], args[1]

		// Loading existing notes from the note.json file
		notes := utils.LoadNotes()

		// Checking if any note with same title exists
		// If it does => updating that note
		for _, m := range notes {
			if m["title"] == title {
				m["body"] = body
				err2 := utils.SaveNotes(notes)
				if err2 != nil {
					return err2
				}
				fmt.Printf("\nUpdated existing note.\n\n")
				return nil
			}
		}

		// Making a new note
		newNote := make(map[string]string)
		newNote["title"] = title
		newNote["body"] = body

		// Appending to existing notes
		notes = append(notes, newNote)

		// Saving updated notes
		err2 := utils.SaveNotes(notes)
		if err2 != nil {
			return err2
		}

		fmt.Println("Note added successfully")

		return nil
	},
}

// Checking if the number of arguments provided are as expected
func addErrFunc(args []string) error {
	if len(args) < 2 || len(args) > 2 {
		return fmt.Errorf("2 arguments expected! title & body")
	}
	return nil
}
