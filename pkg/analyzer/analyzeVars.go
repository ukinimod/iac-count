package analyzer

import (
	"github.com/MaibornWolff/iac-count/pkg/core"
	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func numberOfVars(yamlString string) int {
	return len(input.ReadYamlAsMap(yamlString))
}

func analyzeVarString(yamlString string) map[string]int {
	metrics := analyzeYamlString(yamlString)

	metrics[core.Vars] = numberOfVars(yamlString)
	metrics[core.Complexity] = 1

	return metrics
}

func analyzeVars(path string) map[string]int {
	yamlString := input.ReadFileToString(path)

	return analyzeVarString(yamlString)
}
