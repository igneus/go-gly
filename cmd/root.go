package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gly",
	Short: "translate gly notation to gabc",
	Long: `gly is a custom input format for gregorio, designed
for readability and comfort of the author/transcriber.
This utility translates it to regular gabc.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.UsageFunc()(cmd)
	},
}
