package ansible

import (
	"os"
	"path/filepath"

	"github.com/MaibornWolff/iac-count/internal/util"
	input "github.com/MaibornWolff/iac-count/pkg/input"
	"github.com/MaibornWolff/iac-count/pkg/metrics"
	"github.com/MaibornWolff/iac-count/pkg/model"

	log "github.com/sirupsen/logrus"
)

var calculators = []metrics.MetricCalculator{
	FilesCalculator{},
	LocCalculator{},
	RlocCalculator{},
	CommentsCalculator{},
}

var defaultSkipBaseDirList = [...]string{
	"docs",
	"files",
	"plugins",
	"templates",
}

func calculateMetrics(path string, info os.FileInfo) map[string]metrics.Metric {
	var metricMap = make(map[string]metrics.Metric, len(calculators))

	for _, calc := range calculators {
		var content string

		if !info.IsDir() {
			content = input.ReadFileToString(path)
		}
		metric := calc.Analyze(path, content)
		if metric != nil {
			metricMap[metric.Name()] = metric
		}
	}

	return metricMap
}

func AnalyzeAnsibleProject(root string, skipDirList []string) map[string]model.NodeData {
	fileMetrics := make(map[string]model.NodeData)

	// processing
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// path skipping
			if util.IsHidden(path) {
				log.Infof("Skipping hidden path %s", path)
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				log.Warnf("%s", err)
			}

			basename := filepath.Base(path)

			if info.IsDir() && util.Contains(skipDirList, relativePath) {
				log.Infof("Skipping %s with path %s", basename, relativePath)
				return filepath.SkipDir
			}

			if util.Contains(defaultSkipBaseDirList[:], basename) {
				log.Infof("Skipping %s with path %s", basename, relativePath)
				return filepath.SkipDir
			}

			// analysis
			log.Debugf("Analyzing %s with path %s", basename, relativePath)

			var nodeType NodeType
			if !info.IsDir() {
				nodeType = filetype(relativePath)
			} else {
				nodeType = dirtype(path)
			}

			fileMetrics[relativePath] = model.NodeData{
				Path:     relativePath,
				NodeType: string(nodeType),
				Metrics:  calculateMetrics(path, info),
				Children: make(map[string]string),
			}

			// aggregation
			parentPath := util.ParentPath(relativePath)

			if parent, exists := fileMetrics[parentPath]; exists && parentPath != relativePath {
				parent.Children[basename] = relativePath
				recursiveMetricAggregation(fileMetrics, parentPath, relativePath)
			}

			return nil
		})
	if err != nil {
		log.Warnf("%s", err)
	}

	return fileMetrics
}

func recursiveMetricAggregation(fileMetrics map[string]model.NodeData, parentPath, relativePath string) {
	parentNodeData := fileMetrics[parentPath]
	childNodeData := fileMetrics[relativePath]
	metrics.AggregateMetrics(&(parentNodeData.Metrics), &(childNodeData.Metrics))

	grandparentPath := util.ParentPath(parentPath)
	if _, exists := fileMetrics[grandparentPath]; exists && parentPath != grandparentPath {
		recursiveMetricAggregation(fileMetrics, grandparentPath, relativePath)
	}
}
