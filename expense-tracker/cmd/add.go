package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// type AddCmdFlags struct {
// 	Description string
// 	Amount      int
// }

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Add new expense",
	Long:    "Add a new expense with description and amount",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add new expense babay")
	},
}

func init() {
	cmd := AddCmdFlags{}

	addCmd.Flags().StringVarP(&cmd.Description, "description", "d", "", "Add description to expense")
	addCmd.Flags().IntVarP(&cmd.Amount, "amount", "a", -1, "Add expense amount")
	rootCmd.AddCommand(addExpenseCmd)
}
