package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("Expense Tracker")
	var rootCmd = &cobra.Command{
		Use:   "expense-tracker",
		Short: "CLI tool for managing expenses",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
