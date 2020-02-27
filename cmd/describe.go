package cmd

import (
	"github.com/spf13/cobra"
	"github.com/MaibornWolff/iac-count/pkg/output"
)

var CmdDescribe = &cobra.Command{
	Use:   "describe",
	Short: "`describe` describes the available metrics",
	Long:  "`describe` describes the available metrics",
	Run: func(cmd *cobra.Command, args []string) {
		configureLogging()

		output.PrintMetrics()

	},
}
