package ansible

import (
	"strings"

	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type LocCalculator struct {
}

func (calculator LocCalculator) Analyze(path, content string) metrics.Metric {
	if content == "" {
		return nil
	}

	n := 0
	for _, r := range content {
		if r == '\n' {
			n++
		}
	}
	if len(content) > 0 && !strings.HasSuffix(content, "\n") {
		n++
	}

	return metrics.Loc{n}
}
