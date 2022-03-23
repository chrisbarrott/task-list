package cmd

import (
	"github.com/chrisbarrott/task-list/data"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List out all the tasks.",
	Long:  `List out all the tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayAllTasks()
	},
}

func init() {
	taskCmd.AddCommand(listCmd)
}
