package ansible

import (
	"strings"

	"github.com/MaibornWolff/iac-count/internal/util"
	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type LocCalculator struct {
}

func (calculator LocCalculator) IsFileValidForMetric(path string) bool {
	return util.IsTextFile(path)
}

func (calculator LocCalculator) Analyze(path, content string) metrics.Metric {
	return metrics.Loc{
		Val: strings.Count(content, "\n"),
	}
}
