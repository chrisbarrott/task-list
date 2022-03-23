/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/task-list/data"
)

// updateStatusCmd represents the updateStatus command
var updateStatusCmd = &cobra.Command{
	Use:   "updateStatus",
	Short: "Update the status of a task",
	Long:  `Update the status of a task.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateTaskStatus()
	},
}

func init() {
	taskCmd.AddCommand(updateStatusCmd)
}

func updateTaskStatus() {
	taskPromptContent := PromptContent{
		"Please provide a task name",
		"What task would you like to update? ",
	}

	task := PromptGetInput(taskPromptContent)

	statusPromptContent := PromptContent{
		"Please provide a status",
		fmt.Sprintf("What is the status of this task %s? ", task),
	}

	status := PromptGetSelect(statusPromptContent)

	data.UpdateStatusSQL(task, status)
}
