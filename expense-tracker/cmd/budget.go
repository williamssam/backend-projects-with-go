package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type BudgetCmd struct {
	Month  int
	Amount int
}

var budgetCmd = &cobra.Command{
	Use:   "budget",
	Short: "Add budget for a month",
	Long:  "Add buget for a particular month. If no month is provided, use the current month",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add budget for a month")
	},
}

func init() {
	cmd := BudgetCmd{}

	budgetCmd.Flags().IntVarP(&cmd.Amount, "amount", "a", -1, "Add budget amount for a particular month")
	budgetCmd.Flags().IntVarP(&cmd.Month, "month", "m", -1, "Add budget for a month, if not provided use the current month")
	rootCmd.AddCommand(budgetCmd)
}
