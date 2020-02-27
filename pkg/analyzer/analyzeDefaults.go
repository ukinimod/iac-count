package analyzer

import (
	"github.com/MaibornWolff/iac-count/pkg/core"
	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func numberOfDefaults(yamlString string) int {
	return len(input.ReadYamlAsMap(yamlString))
}

func analyzeDefaultsString(yamlString string) map[string]int {
	metrics := analyzeYamlString(yamlString)

	metrics[core.Defaults] = numberOfDefaults(yamlString)
	metrics[core.Complexity] = 1

	return metrics
}

func analyzeDefaults(path string) map[string]int {
	yamlString := input.ReadFileToString(path)

	return analyzeDefaultsString(yamlString)
}
