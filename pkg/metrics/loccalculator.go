package metrics

import (
	"strings"
)

type LocCalculator struct {
}

func (calculator LocCalculator) Analyze(path, content string) Metric {
	return Loc{
		Val: strings.Count(content, "\n"),
	}
}
