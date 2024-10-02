package cmd

import (
	"fmt"

	"github.com/fhasnur/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
)

var month int

func NewSummaryCmd() *cobra.Command {
	summaryCmd := &cobra.Command{
		Use:   "summary",
		Short: "Show the total expenses",
		Run: func(cmd *cobra.Command, args []string) {
			err := expense.SummaryExpenses(month)
			if err != nil {
				fmt.Println("Error:", err)
			}
		},
	}

	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "Month for summary")

	return summaryCmd
}
