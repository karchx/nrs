package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nrs",
	Short: "A simple tool that collects TODOs in the source code and reports them as GitHub issues.",
	Long: `Language agnostic tool that collects TODOs in the source code and reports them as Issues`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {	
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


