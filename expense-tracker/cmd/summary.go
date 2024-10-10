package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type SummaryCmd struct {
	Month int
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary of all expenses",
	Long:  "Gives total amount of all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add new expense babay")
	},
}

func init() {
	cmd := SummaryCmd{}

	summaryCmd.Flags().IntVarP(&cmd.Month, "month", "m", -1, "Get summary for a particular month")
	rootCmd.AddCommand(summaryCmd)

}
