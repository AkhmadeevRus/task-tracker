package cmd

import (
	"errors"

	"github.com/AkhmadeevRus/task-tracker/pkg/models"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a task to the task list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}
	return cmd
}

func RunAddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}
	return models.AddTask(args[0])
}
