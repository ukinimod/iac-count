package analyzer

import (
	"github.com/MaibornWolff/iac-count/pkg/core"
	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func numberOfHandlers(handlerYaml string) int {
	return len(input.ReadHandlersString(handlerYaml))
}

func analyzeHandler(path string) map[string]int {
	yamlString := input.ReadFileToString(path)

	metrics := analyzeYamlString(yamlString)

	metrics[core.Handlers] = numberOfHandlers(yamlString)
	metrics[core.Complexity] = metrics[core.Handlers]

	return metrics
}
