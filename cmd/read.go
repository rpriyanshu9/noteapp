package cmd

import (
	"fmt"

	"github.com/rpriyanshu9/noteapp/utils"
	"github.com/spf13/cobra"
)

var readField = ""

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVar(&readField, "by", "", "read note with given field type")
	readCmd.MarkFlagRequired("by")
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a note",
	Long:  "\nReads a note for the title given by user",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Checking for any errors. Specifically the "too few/too many arguments" error
		// if err := readErrFunction(args); err != nil {
		// 	return err
		// }

		if readField == "" {
			fmt.Printf("\nNo field provided.\n\n")
			return nil
		}

		if len(args) != 1 {
			fmt.Printf("\nNo field value provided.\n\n")
			return nil
		}

		fieldValue := args[0]

		// Getting title from arguments provided
		// title := args[0]

		// Loading existing notes from the note.json file
		notes := utils.LoadNotes()

		// Searching for the note which matches the title
		for _, m := range notes {
			switch readField {
			case "title", "Title":
				if m.Title == fieldValue {
					fmt.Printf("\nYour note is: \n-> %s: %s\n\n", m.Title, m.Body)
					return nil
				}
			case "body", "Body":
				if m.Body == fieldValue {
					fmt.Printf("\nYour note is: \n-> %s: %s\n\n", m.Title, m.Body)
					return nil
				}
			}
		}

		// If we reach here, note with the provided title doesn't exist
		fmt.Printf("\nError: no note with %s: \"%s\" present\n\n", readField, fieldValue)
		return nil
	},
}

// Checking if the number of arguments provided are as expected
// func _(args []string) error {
// 	if len(args) < 1 || len(args) > 1 {
// 		return fmt.Errorf("1 argument expected")
// 	}
// 	return nil
// }
