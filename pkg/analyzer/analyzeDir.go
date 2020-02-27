package analyzer

import "github.com/MaibornWolff/iac-count/pkg/core"

func analyzeDir(path string) map[string]int {
	metrics := make(map[string]int)

	metrics[core.Files] = recursiveFileCount(path)

	return metrics
}
