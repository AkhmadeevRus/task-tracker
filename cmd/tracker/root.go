package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tarcker",
		Short: "Task Tracker is a CLI tool for managing tasks",
	}
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewUpdateCmd())
	return cmd
}
