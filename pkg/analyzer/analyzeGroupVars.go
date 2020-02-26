package analyzer

import (
	"github.com/ukinimod/iac-count/pkg/core"
	input "github.com/ukinimod/iac-count/pkg/input"
)

func numberOfGroupVars(yamlString string) int {
	return len(input.ReadYamlAsMap(yamlString))
}

func analyzeGroupVarsString(yamlString string) map[string]int {
	metrics := analyzeYamlString(yamlString)

	metrics[core.GroupVars] = numberOfGroupVars(yamlString)
	metrics[core.Complexity] = 1

	return metrics
}

func analyzeGroupVars(path string) map[string]int {
	yamlString := input.ReadFileToString(path)

	return analyzeGroupVarsString(yamlString)
}
