package ansible

import "github.com/MaibornWolff/iac-count/pkg/metrics"

type PlaysCalculator struct {
}

func (calculator PlaysCalculator) analyze(path, content string) metrics.Metric {
	return metrics.Plays{len(ReadPlaybookString(content))}
}
