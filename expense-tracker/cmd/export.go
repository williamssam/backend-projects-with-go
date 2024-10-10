package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export expenses to CSV",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Export all expenses to CSV")
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
