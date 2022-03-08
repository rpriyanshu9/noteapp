package cmd

import (
	"fmt"

	"github.com/rpriyanshu9/noteapp/utils"
	"github.com/spf13/cobra"
)

var removeField = ""

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringVar(&removeField, "by", "", "field name to search and delete")
	removeCmd.MarkFlagRequired("by")
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a note",
	Long:  "Removes a note with the title given by the user",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Checking for any errors. Specifically the "too few/too many arguments" error
		// if err := removeErrFunc(args); err != nil {
		// 	return err
		// }

		if removeField == "" {
			fmt.Printf("\nNo field provided.\n\n")
			return nil
		}

		if len(args) != 1 {
			fmt.Printf("\nNo field value provided.\n\n")
			return nil
		}

		fieldValue := args[0]

		// Loading existing notes from the note.json file
		notes := utils.LoadNotes()

		// Temporary empty list of notes
		result := []utils.Note{}

		// Appending every non matching note to this new result slice
		for _, m := range notes {
			switch removeField {
			case "title", "Title":
				if m.Title != fieldValue {
					result = append(result, m)
				}
			case "body", "Body":
				if m.Body != fieldValue {
					result = append(result, m)
				}
			default:
				result = append(result, m)
			}
		}

		// Saving updated notes
		err := utils.SaveNotes(result)
		if err != nil {
			return err
		}

		fmt.Println("Note deleted successfully")

		return nil
	},
}

// func _(args []string) error {
// 	if len(args) < 1 || len(args) > 1 {
// 		return fmt.Errorf("1 argument expected")
// 	}
// 	return nil
// }
