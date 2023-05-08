package cmd

import (
	"github.com/karchx/nrs/pkg/errors"
	"github.com/karchx/nrs/pkg/todo"
	"github.com/karchx/nrs/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos of a dir recrusively",
	Run: func(cmd *cobra.Command, args []string) {
		project := utils.GetProject(".")

		reported, _ := cmd.Flags().GetBool("reported")
		unreported, _ := cmd.Flags().GetBool("reported")

		err := utils.ListSubCommand(*project, func(todoP todo.Todo) bool {
			filter := reported == unreported

			if unreported {
				filter = filter || todoP.ID == nil
			}

			if reported {
				filter = filter || todoP.ID != nil
			}

			return filter
		})

		errors.ExitOnError(err)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().Bool("unreported", false, "Unreported")
	listCmd.PersistentFlags().Bool("reported", false, "Reported")
}
