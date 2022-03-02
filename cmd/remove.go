package cmd

import (
	"fmt"

	"github.com/rpriyanshu9/noteapp/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove title",
	Short: "Removes a note",
	Long:  "Removes a note with the title given by the user",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := removeErrFunc(args); err != nil {
			return err
		}
		title := args[0]
		// fmt.Println(title)
		notes := utils.LoadNotes()
		result := []map[string]string{}
		for _, m := range notes {
			if m["title"] != title {
				result = append(result, m)
			}
		}
		err2 := utils.SaveNotes(result)
		if err2 != nil {
			return err2
		}
		fmt.Println("Note deleted successfully")
		return nil
	},
}

func removeErrFunc(args []string) error {
	if len(args) < 1 || len(args) > 1 {
		return fmt.Errorf("1 argument expected")
	}
	return nil
}
