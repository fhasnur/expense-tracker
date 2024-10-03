package cmd

import (
	"fmt"
	"strings"

	"github.com/fhasnur/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all expenses",
		Run: func(cmd *cobra.Command, args []string) {
			expenses, err := expense.ListExpenses()
			if err != nil {
				fmt.Println("Error listing expenses:", err)
			}

			fmt.Printf(
				"%-3s | %-20s | %-15s | %-10s\n",
				"ID",
				"Date/Time",
				"Description",
				"Amount",
			)
			fmt.Println(strings.Repeat("-", 55))

			for _, expense := range expenses {
				fmt.Printf(
					"%-3d | %-20s | %-15s | $%-10.0f\n",
					expense.ID,
					expense.UpdatedAt.Format("2006-01-02 15:04"),
					expense.Description,
					expense.Amount,
				)
			}
		},
	}
	return listCmd
}
