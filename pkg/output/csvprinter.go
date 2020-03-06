package output

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/MaibornWolff/iac-count/internal/util"
	"github.com/MaibornWolff/iac-count/pkg/model"
)

const (
	PrintLevelRole    = "role"
	PrintLevelFile    = "file"
	PrintLevelProject = "project"
)

func csvHeader(metricNames []string) string {
	return "path,type," + strings.Join(metricNames, ",")
}

func csvBodyLine(k string, v model.NodeData, metricNames []string) string {
	var sb strings.Builder
	sb.WriteString(k)
	sb.WriteString(",")
	sb.WriteString(v.NodeType)
	for _, metricName := range metricNames {
		sb.WriteString(",")
		for _, metric := range v.Metrics {
			if metric != nil && metric.Name() == metricName {
				sb.WriteString(strconv.Itoa(metric.Value()))
			}
		}
	}

	return sb.String()
}

func PrintMetricsAsCsv(metrics map[string]model.NodeData, level string) {
	switch level {
	case PrintLevelFile:
		printAsCsv(metrics, func(it string) bool { return true })
	case PrintLevelRole:
		printAsCsv(metrics, func(it string) bool { return it == "role" || it == "ansible_project" })
	case PrintLevelProject:
		printAsCsv(metrics, func(it string) bool { return it == "ansible_project" })
	default:
		log.Fatalf("Unknown printing level: %s", level)
	}
}

func calculatedMetricNames(metrics map[string]model.NodeData) []string {
	var metricNames []string

	for _, node := range metrics {
		for metricName := range node.Metrics {
			if !util.Contains(metricNames, metricName) {
				metricNames = append(metricNames, metricName)
				sort.Strings(metricNames)
			}
		}
	}

	return metricNames
}

func printAsCsv(metrics map[string]model.NodeData, filter func(string) bool) {
	metricNames := calculatedMetricNames(metrics)
	fmt.Println(csvHeader(metricNames))

	keys := make([]string, 0, len(metrics))
	for k := range metrics {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if filter(metrics[k].NodeType) {
			fmt.Println(csvBodyLine(k, metrics[k], metricNames))
		}
	}
}
