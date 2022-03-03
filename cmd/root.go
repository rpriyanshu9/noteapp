package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "noteapp",
		Short: "A very basic CLI tool for taking notes",
		Long:  `A very basic CLI tool that implements CRUD operations, written in Golang`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welocme to noteapp!")
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
