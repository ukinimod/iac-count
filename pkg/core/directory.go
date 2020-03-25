package core

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/MaibornWolff/iac-count/internal/util"
	"github.com/MaibornWolff/iac-count/pkg/metrics"
	log "github.com/sirupsen/logrus"
)

type Directory struct {
	path string
}

func (subject Directory) Path() string {
	return subject.path
}

func (subject Directory) Analyze() map[string]metrics.Node {
	nodes := make(map[string]metrics.Node, 1)

	fileinfo, err := ioutil.ReadDir(subject.path)
	if err == nil {
		for i := range fileinfo {
			info := fileinfo[i]
			path := filepath.Join(subject.path, info.Name())
			if util.IsHidden(path) {
				log.Infof("Skipping hidden path %s", path)
				continue
			}

			log.Debugf("Analyzing %s", path)

			subject := createSubject(path, info)

			for k, v := range subject.Analyze() {
				nodes[k] = v
			}
		}
	} else {
		log.Errorf("%s", err)
	}

	nodes[subject.path] = metrics.Node{
		Path:     subject.path,
		NodeType: "dir",
		Metrics:  subject.CalculateMetrics(),
	}

	return nodes
}

var fileCreator = FileCreator{}
var yamlfileCreator = YamlfileCreator{}
var textfileCreator = TextfileCreator{}
var directoryCreator = DirectoryCreator{}

func createSubject(path string, info os.FileInfo) Subject {
	if !info.IsDir() {
		if util.IsYamlFile(path) {
			return yamlfileCreator.CreateFromPath(path, info)
		} else if util.IsTextFile(path) {
			return textfileCreator.CreateFromPath(path, info)
		} else {
			return fileCreator.CreateFromPath(path, info)
		}
	} else {
		return directoryCreator.CreateFromPath(path, info)
	}
}

func (subject Directory) CalculateMetrics() map[string]metrics.Metric {
	path := subject.Path()
	var metricMap = make(map[string]metrics.Metric, len(directoryCalculators))

	for _, calc := range directoryCalculators {
		metric := calc.Analyze(path, "")
		metricMap[metric.Name()] = metric
	}

	return metricMap
}

var directoryCalculators = []metrics.MetricCalculator{
	metrics.FilesCalculator{},
}

type DirectoryCreator struct {
}

func (creator DirectoryCreator) CreateFromPath(path string, info os.FileInfo) Subject {
	return Directory{
		path: path,
	}
}
