package cmd

import (
	"fmt"

	"github.com/chrisbarrott/task-list/data"
	"github.com/spf13/cobra"
)

// updateStatusCmd represents the updateStatus command
var updateStatusCmd = &cobra.Command{
	Use:   "updateStatus",
	Short: "Update the status of a task",
	Long:  `Update the status of a task.`,
	Run: func(cmd *cobra.Command, args []string) {
		UpdateTaskStatus()
	},
}

func init() {
	taskCmd.AddCommand(updateStatusCmd)
}

func UpdateTaskStatus() {
	// task update prompt
	taskPromptContent := PromptContent{
		"Please provide a task name",
		"What task would you like to update? ",
	}
	task := PromptGetInput(taskPromptContent)

	// new status prompt
	statusPromptContent := PromptContent{
		"Please provide a status",
		fmt.Sprintf("What is the status of this task %s? ", task),
	}
	status := PromptGetSelect(statusPromptContent)

	// update SQL
	data.UpdateStatusSQL(task, status)
}
