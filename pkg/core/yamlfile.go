package core

import (
	"os"

	input "github.com/MaibornWolff/iac-count/pkg/input"

	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type Yamlfile struct {
	path string
}

func (subject Yamlfile) Path() string {
	return subject.path
}

func (subject Yamlfile) Analyze() map[string]metrics.Node {
	nodes := make(map[string]metrics.Node, 1)
	nodes[subject.path] = metrics.Node{
		Path:     subject.path,
		NodeType: "yaml",
		Metrics:  subject.CalculateMetrics(),
	}

	return nodes
}

var yamlfileCalculators = []metrics.MetricCalculator{
	metrics.LocCalculator{},
	metrics.RlocCalculator{},
	metrics.CommentlinesCalculator{},
}

func (subject Yamlfile) CalculateMetrics() map[string]metrics.Metric {
	path := subject.Path()
	var metricMap = make(map[string]metrics.Metric, len(fileCalculators))

	for _, calc := range yamlfileCalculators {
		content := input.ReadFileToString(path)
		metric := calc.Analyze(path, content)
		metricMap[metric.Name()] = metric
	}

	return metricMap
}

type YamlfileCreator struct {
}

func (creator YamlfileCreator) CreateFromPath(path string, info os.FileInfo) Subject {
	return Yamlfile{
		path: path,
	}
}
