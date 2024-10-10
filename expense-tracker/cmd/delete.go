package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type DeleteCmdFlags struct {
	Id int
}

var deleteCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"delete"},
	Short:   "Delete expense",
	Long:    "Delete exisiting expense",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete expense")
	},
}

func init() {
	cmd := DeleteCmdFlags{}

	deleteCmd.Flags().IntVarP(&cmd.Id, "id", "", -1, "ID of expense to delete")
	rootCmd.AddCommand(deleteCmd)
}
