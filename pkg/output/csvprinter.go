package output

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	core "github.com/MaibornWolff/iac-count/pkg/core"
)

func csvHeader(metricNames []string) string {
	return "path,type," + strings.Join(metricNames, ",")
}

func csvBodyLine(k string, v core.FileData, metricNames []string) string {
	var sb strings.Builder
	sb.WriteString(k)
	sb.WriteString(",")
	sb.WriteString(v.Filetype)
	for _, metric := range metricNames {
		sb.WriteString(",")
		if metricValue, ok := v.Metrics[metric]; ok {
			sb.WriteString(strconv.Itoa(metricValue))
		}
	}

	return sb.String()
}

func PrintMetricsAsCsv(metrics map[string]core.FileData, metricNames []string) {
	printAsCsv(metrics, metricNames, func(it string) bool { return true })
}

func PrintRolesAsCsv(metrics map[string]core.FileData, metricNames []string) {
	printAsCsv(metrics, metricNames, func(it string) bool { return it == "role" || it == "ansible_project" })
}

func PrintProjectAsCsv(metrics map[string]core.FileData, metricNames []string) {
	printAsCsv(metrics, metricNames, func(it string) bool { return it == "ansible_project" })
}

func printAsCsv(metrics map[string]core.FileData, metricNames []string, filter func(string) bool) {
	fmt.Println(csvHeader(metricNames))

	keys := make([]string, 0, len(metrics))
	for k := range metrics {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if filter(metrics[k].Filetype) {
			fmt.Println(csvBodyLine(k, metrics[k], metricNames))
		}
	}
}

func PrintMetrics() {
	fmt.Println("metric,description")
	for _, metric := range core.MetricNames {
		fmt.Printf("%s,\"%s\"\n", metric, core.DescribeMetric(metric))
	}
}
