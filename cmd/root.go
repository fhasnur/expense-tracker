package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expense-tracker",
		Short: "Expense Tracker is a CLI tool for managing expenses.",
		Long:  "Manage your expenses with ease using Expense Tracker",
	}

	cmd.AddCommand(
		NewAddCmd(),
		NewListCmd(),
		NewDeleteCmd(),
		NewSummaryCmd(),
	)

	return cmd
}

func Execute() error {
	return NewRootCmd().Execute()
}
