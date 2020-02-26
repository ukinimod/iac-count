package analyzer

import "github.com/ukinimod/iac-count/pkg/core"

func analyzeDir(path string) map[string]int {
	metrics := make(map[string]int)

	metrics[core.Files] = recursiveFileCount(path)

	return metrics
}
