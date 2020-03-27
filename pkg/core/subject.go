package core

import (
	"os"

	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type Subject interface {
	Path() string
	CalculateMetrics() map[string]metrics.Metric
	Analyze() map[string]metrics.Node
}

type SubjectCreator interface {
	CreateFromPath(path string, info os.FileInfo) Subject
}
