package cmd

import (
	"fmt"

	"github.com/fhasnur/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
)

var id int64

func NewDeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an expense by ID",
		Run: func(cmd *cobra.Command, args []string) {
			err := expense.DeleteExpense(id)
			if err != nil {
				fmt.Println("Error deleting expense:", err)
			} else {
				fmt.Println("Expense deleted successfully")
			}
		},
	}

	deleteCmd.Flags().Int64VarP(&id, "id", "i", 0, "ID of the expense to delete (required)")
	deleteCmd.MarkFlagRequired("id")

	return deleteCmd
}
