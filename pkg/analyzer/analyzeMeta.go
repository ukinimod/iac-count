package analyzer

import (
	"github.com/MaibornWolff/iac-count/pkg/core"
	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func numberOfDependencies(yamlString string) int {
	return len(input.ReadMetaString(yamlString).AnsibleDependencies)
}

func analyzeMeta(path string) map[string]int {
	yamlString := input.ReadFileToString(path)

	metrics := analyzeYamlString(yamlString)

	metrics[core.Dependencies] = numberOfDependencies(yamlString)
	metrics[core.Complexity] = 1

	return metrics
}
