package cmd

import (
	"github.com/igneus/go-gly/gly"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	RootCmd.AddCommand(gabcCmd)
}

var gabcCmd = &cobra.Command{
	Use:   "gabc",
	Short: "Convert gly to gabc",
	Long:  `Expects gly document, produces one or more gabc files.`,
	Run: func(cmd *cobra.Command, args []string) {
		var parser gly.Parser
		parser.Parse(os.Stdin)
	},
}
