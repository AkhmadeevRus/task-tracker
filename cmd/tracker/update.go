package cmd

import (
	"errors"
	"strconv"

	"github.com/AkhmadeevRus/task-tracker/pkg/models"
	"github.com/spf13/cobra"
)

func NewUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "update task by id and description",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateTaskCmd(args)
		},
	}
	return cmd
}

func RunUpdateTaskCmd(args []string) error {
	if len(args) != 2 {
		return errors.New("please provide a task id and new description")
	}

	taskId, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return err
	}

	newDescription := args[1]
	return models.UpdateTaskDescription(taskId, newDescription)
}
