package cmd

import (
	"log"

	"github.com/spf13/cobra"
	analyzer "github.com/ukinimod/iac-count/pkg/analyzer"
	"github.com/ukinimod/iac-count/pkg/core"
	"github.com/ukinimod/iac-count/pkg/output"
)

func getValidMetrics() []string {
	var metricList []string
	for _, metric := range MetricList {
		if core.IsValidMetric(metric) {
			metricList = append(metricList, metric)
		} else {
			log.Printf("[WARN] Metric %s is not a valid metric, see `describe` command for the list of valid metrics.\n", metric)
		}
	}

	return metricList
}

var CmdAnsible = &cobra.Command{
	Use:   "ansible [root of ansible project]",
	Short: "`ansible` analyzes ansible projects",
	Long:  "`ansible` analyzes ansible projects",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configureLogging()

		metricList := getValidMetrics()

		fileMetrics := analyzer.AnalyzeAnsibleProject(args[0])

		switch PrintLevel {
		case "file":
			output.PrintMetricsAsCsv(fileMetrics, metricList)
		case "role":
			output.PrintRolesAsCsv(fileMetrics, metricList)
		case "project":
			output.PrintProjectAsCsv(fileMetrics, metricList)
		default:
			log.Fatalf("Unknown printing level: %s", PrintLevel)
		}

	},
}
