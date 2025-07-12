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

func NewStatusInProgressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Mark-in-progress",
		Short: "change task status to 'in-progress'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, models.TASK_STATUS_IN_PROGRESS)
		},
	}
	return cmd
}

func NewStatusTodoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Mark-i-progress",
		Short: "change task status to 'in-progress'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, models.TASK_STATUS_TODO)
		},
	}
	return cmd
}

func NewStatusDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Mark-done",
		Short: "change task status to 'done'",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateStatusCmd(args, models.TASK_STATUS_DONE)
		},
	}
	return cmd
}

func RunUpdateStatusCmd(args []string, status models.TaskStatus) error {
	if len(args) == 0 {
		return errors.New("task id is required")
	}

	id, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return err
	}

	return models.UpdateTaskStatus(id, status)
}
