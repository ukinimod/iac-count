package core

import (
	"os"

	input "github.com/MaibornWolff/iac-count/pkg/input"

	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type File struct {
	path string
}

func (subject File) Path() string {
	return subject.path
}

func (subject File) Analyze() map[string]metrics.Node {
	nodes := make(map[string]metrics.Node, 1)
	nodes[subject.path] = metrics.Node{
		Path:     subject.path,
		NodeType: "file",
		Metrics:  subject.CalculateMetrics(),
	}

	return nodes
}

var fileCalculators = []metrics.MetricCalculator{}

func (subject File) CalculateMetrics() map[string]metrics.Metric {
	path := subject.Path()
	var metricMap = make(map[string]metrics.Metric, len(fileCalculators))

	for _, calc := range fileCalculators {
		content := input.ReadFileToString(path)
		metric := calc.Analyze(path, content)
		metricMap[metric.Name()] = metric
	}

	return metricMap
}

type FileCreator struct {
}

func (creator FileCreator) CreateFromPath(path string, info os.FileInfo) Subject {
	return File{
		path: path,
	}
}
