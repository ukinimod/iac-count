package core

import (
	"os"

	input "github.com/MaibornWolff/iac-count/pkg/input"

	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type Textfile struct {
	path string
}

func (subject Textfile) Path() string {
	return subject.path
}

func (subject Textfile) Analyze() map[string]metrics.Node {
	nodes := make(map[string]metrics.Node, 1)
	nodes[subject.path] = metrics.Node{
		Path:     subject.path,
		NodeType: "file",
		Metrics:  subject.CalculateMetrics(),
	}

	return nodes
}

var textfileCalculators = []metrics.MetricCalculator{
	metrics.LocCalculator{},
}

func (subject Textfile) CalculateMetrics() map[string]metrics.Metric {
	path := subject.Path()
	var metricMap = make(map[string]metrics.Metric, len(fileCalculators))

	for _, calc := range textfileCalculators {
		content := input.ReadFileToString(path)
		metric := calc.Analyze(path, content)
		metricMap[metric.Name()] = metric
	}

	return metricMap
}

type TextfileCreator struct {
}

func (creator TextfileCreator) CreateFromPath(path string, info os.FileInfo) Subject {
	return Textfile{
		path: path,
	}
}
