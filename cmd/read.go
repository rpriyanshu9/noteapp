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
		if err := readErrFunction(args); err != nil {
			return err
		}
		title := args[0]
		notes := utils.LoadNotes()
		for _, m := range notes {
			if m["title"] == title {
				fmt.Printf("\nYour note is: \n-> %s: %s\n\n", m["title"], m["body"])
				return nil
			}
		}
		fmt.Printf("\nError: no note with title: \"%s\" present\n\n", title)
		return nil
	},
}

func readErrFunction(args []string) error {
	if len(args) < 1 || len(args) > 1 {
		return fmt.Errorf("1 argument expected")
	}
	return nil
}
