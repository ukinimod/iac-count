package ansible

import (
	reader "github.com/MaibornWolff/iac-count/pkg/input"
	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type VarsCalculator struct {
}

func (calculator VarsCalculator) Analyze(path, content string) metrics.Metric {
	return metrics.Vars{len(reader.ReadYamlAsMap(content))}
}
