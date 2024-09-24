package cmd

import (
	"fmt"

	"github.com/fhasnur/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
)

var description string
var amount float64

func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new expense",
		Run: func(cmd *cobra.Command, args []string) {
			if amount < 0 {
				fmt.Println("Error: The amount cannot be negative")
				return
			}

			id, err := expense.AddExpense(description, amount)
			if err != nil {
				fmt.Println("Error adding expense:", err)
			} else {
				fmt.Printf("Expense added successfully (ID: %d)\n", id)
			}
		},
	}

	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the expense (required)")
	addCmd.MarkFlagRequired("description")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Amount of the expense (required)")
	addCmd.MarkFlagRequired("amount")

	return addCmd
}
