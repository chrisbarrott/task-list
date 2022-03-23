package cmd

import (
	"github.com/chrisbarrott/task-list/data"
	"github.com/spf13/cobra"
)

// completeTaskCmd represents the completeTask command
var completeTaskCmd = &cobra.Command{
	Use:   "completeTask",
	Short: "Complete and remove task",
	Long:  `Complete and remove task.`,
	Run: func(cmd *cobra.Command, args []string) {
		removeCompletedTask()
	},
}

func init() {
	taskCmd.AddCommand(completeTaskCmd)
}

func removeCompletedTask() {
	taskPromptContent := PromptContent{
		"Please provide a task name",
		"What task would you like to delete? ",
	}

	task := PromptGetInput(taskPromptContent)
	/*
		statusPromptContent := PromptContent{
			"Please provide a status",
			fmt.Sprintf("What is the status of this task %s?", task),
		}

		status := PromptGetSelect(statusPromptContent)
	*/
	data.RemoveTaskSQL(task)
}
