/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/task-list/data"
)

type PromptContent struct {
	errorMsg string
	label    string
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new task",
	Long:  `Creates a new task.`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewTask()
	},
}

func init() {
	taskCmd.AddCommand(newCmd)

}

func PromptGetInput(pc PromptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func PromptGetSelect(pc PromptContent) string {
	items := []string{"Backlog", "In Progress", "Completed"}
	index := -1

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result

}

func createNewTask() {
	taskPromptContent := PromptContent{
		"Please provide a task",
		"What task would you like to create? ",
	}

	task := PromptGetInput(taskPromptContent)

	statusPromptContent := PromptContent{
		"Please provide a status",
		fmt.Sprintf("What is the status of this task %s? ", task),
	}

	status := PromptGetSelect(statusPromptContent)

	// get the unix time of now
	createdAt := strconv.Itoa(int(time.Now().Unix()))

	// format the int
	i, err := strconv.ParseInt(createdAt, 10, 64)
	if err != nil {
		panic(err)
	}

	// conver the UTC time to string because thats how its declared in the database
	tm := time.Unix(i, 0).String()
	fmt.Println("Task created at: ", tm)

	// execute insert
	data.InsertTask(task, status, tm)
}
