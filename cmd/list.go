package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos of a dir recrusively",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().Bool("unreported", false, "Unreported")
	listCmd.PersistentFlags().Bool("reported", false, "Reported")
}
