package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(gabcCmd)
}

var gabcCmd = &cobra.Command{
	Use:   "gabc",
	Short: "Convert gly go gabc",
	Long:  `Expects gly document, produces one or more gabc files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hey, happily here. But it doesn't do anything yet.")
	},
}
