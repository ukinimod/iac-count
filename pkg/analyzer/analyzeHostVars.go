package analyzer

import (
	"github.com/ukinimod/iac-count/pkg/core"
	input "github.com/ukinimod/iac-count/pkg/input"
)

func numberOfHostVars(yamlString string) int {
	return len(input.ReadYamlAsMap(yamlString))
}

func analyzeHostVarsString(yamlString string) map[string]int {
	metrics := analyzeYamlString(yamlString)

	metrics[core.HostVars] = numberOfHostVars(yamlString)
	metrics[core.Complexity] = 1

	return metrics
}
func analyzeHostVars(path string) map[string]int {
	yamlString := input.ReadFileToString(path)

	return analyzeHostVarsString(yamlString)
}
