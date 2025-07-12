package cmd

import (
	"github.com/AkhmadeevRus/task-tracker/pkg/models"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTaskCmd(args)
		},
	}
	return cmd
}

func RunListTaskCmd(args []string) error {
	if len(args) > 0 {
		status := models.TaskStatus(args[0])
		return models.ListTasks(status)
	}
	return models.ListTasks("all")
}
