package cmd

import (
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "A task is any task to do or complete",
	Long: `The options here are to:
	- Create new task
	- List all tasks
	- Update a task 
	- complete (remove) a task.`,
	// Removed Run to just run the description
}

func init() {
	rootCmd.AddCommand(taskCmd)
}
