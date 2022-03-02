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
		if err := addErrFunc(args); err != nil {
			return err
		}
		title, body := args[0], args[1]
		notes := utils.LoadNotes()
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
		newNote := make(map[string]string)
		newNote["title"] = title
		newNote["body"] = body
		notes = append(notes, newNote)
		err2 := utils.SaveNotes(notes)
		if err2 != nil {
			return err2
		}
		fmt.Println("Note added successfully")
		return nil
	},
}

func addErrFunc(args []string) error {
	if len(args) < 2 || len(args) > 2 {
		return fmt.Errorf("2 arguments expected! title & body")
	}
	return nil
}
