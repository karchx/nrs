package cmd

import (
	log "github.com/gothew/l-og"
	"github.com/spf13/cobra"
)

var (
	prependBody string

	reportCmd = &cobra.Command{
		Use:   "report",
		Short: "Report open issues github",
		Run: func(cmd *cobra.Command, args []string) {
      log.Info(prependBody)
		},
	}
)

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.PersistentFlags().StringVar(&prependBody, "prepend-body", "", "Prepend body")
}
