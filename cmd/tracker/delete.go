package cmd

import (
	"errors"
	"strconv"

	"github.com/AkhmadeevRus/task-tracker/pkg/models"
	"github.com/spf13/cobra"
)

func NewDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a task by the passed id",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteTaskCmd(args)
		},
	}
	return cmd
}

func RunDeleteTaskCmd(args []string) error {
	if len(args) != 1 {
		return errors.New("please provide a task ID")
	}

	taskID, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return err
	}
	return models.DeleteTaskById(taskID)
}
