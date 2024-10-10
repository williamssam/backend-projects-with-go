package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type AddCmdFlags struct {
	Description string
	Amount      int
	Category    string
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"addition"},
	Short:   "Add new expense",
	Long:    "Add a new expense with description and amount",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add new expense babay")
	},
}

func init() {
	cmd := AddCmdFlags{}

	addCmd.Flags().StringVarP(&cmd.Description, "description", "d", "", "Add description to expense")
	addCmd.Flags().IntVarP(&cmd.Amount, "amount", "a", -1, "Add expense amount")
	addCmd.Flags().StringVarP(&cmd.Category, "category", "c", "", "Add category to expense")
	rootCmd.AddCommand(addCmd)
}
