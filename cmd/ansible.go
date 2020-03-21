package cmd

import (
	analyzer "github.com/MaibornWolff/iac-count/pkg/ansible"
	"github.com/MaibornWolff/iac-count/pkg/output"
	"github.com/spf13/cobra"
)

var CmdAnsible = &cobra.Command{
	Use:   "ansible ANSIBLE_ROOT",
	Short: "`ansible` analyzes ansible projects",
	Long:  "`ansible` analyzes ansible projects",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configureLogging()

		skipDirList := make([]string, 0) // TODO: Extension for future feature
		fileMetrics := analyzer.AnalyzeAnsibleProject(args[0], skipDirList)
		output.PrintMetricsAsCsv(fileMetrics, PrintLevel)
	},
}
