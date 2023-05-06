package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos of a dir recrusively",
	Run: func(cmd *cobra.Command, args []string) {
    reported, _ := cmd.Flags().GetBool("reported")
    unreported, _ := cmd.Flags().GetBool("reported")
    fmt.Println(reported, unreported)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().Bool("unreported", false, "Unreported")
	listCmd.PersistentFlags().Bool("reported", false, "Reported")
}
