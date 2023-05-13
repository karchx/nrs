package cmd

import (
	"github.com/karchx/nrs/pkg/platform"
	"github.com/spf13/cobra"
)

var (
	prependBody string

	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Report open issues github",
		Run: func(cmd *cobra.Command, args []string) {
			platform.GetCredentials()
		},
	}
)

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.PersistentFlags().StringVar(&prependBody, "prepend-body", "", "Prepend body")
}
