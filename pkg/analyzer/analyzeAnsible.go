package analyzer

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	core "github.com/ukinimod/iac-count/pkg/core"
)

var analyzers map[string]interface{} = map[string]interface{}{
	"task":            analyzeTask,
	"handler":         analyzeHandler,
	"vars":            analyzeVars,
	"defaults":        analyzeDefaults,
	"meta":            analyzeMeta,
	"host_vars":       analyzeHostVars,
	"group_vars":      analyzeGroupVars,
	"playbook":        analyzePlaybook,
	"file":            nil,
	"yml":             analyzeYaml,
	"role":            analyzeRole,
	"ansible_project": analyzeAnsibleDir,
	"dir":             analyzeDir,
}

var aggregationMetrics = [...]string{
	core.Loc,
	core.Tasks,
	core.Handlers,
	core.Complexity,
	core.Vars,
	core.Defaults,
	core.GroupVars,
	core.HostVars,
	core.Dependencies,
	core.Rloc,
	core.CommentLines,
	core.Templates,
	core.StaticFiles,
	core.Plugins,
	core.Roles,
	core.Plays,
	core.CustomFacts,
	core.TaggedTasks,
}

func AnalyzeAnsibleProject(root string) map[string]core.FileData {
	fileMetrics := make(map[string]core.FileData)

	// processing
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			relativePath, err := filepath.Rel(root, path)
			if err != nil {
				log.Printf("[WARN] %s", err)
			}

			log.Printf("[DEBUG] Analyzing path %s", relativePath)

			basename := filepath.Base(path)
			parentPath := strings.TrimSuffix(relativePath, string(filepath.Separator)+basename)
			if parentPath == "" || parentPath == relativePath {
				parentPath = "."
			}

			var typeOfNode string
			if !info.IsDir() {
				typeOfNode = filetype(relativePath)
			} else {
				typeOfNode = dirtype(path)

				if basename == "files" || basename == "templates" || basename == "plugins" || basename == "docs" {
					return filepath.SkipDir
				}
			}

			analyzer := analyzers[typeOfNode]
			if analyzer != nil {

				fileMetrics[relativePath] = core.FileData{
					RelativePath: relativePath,
					Filetype:     typeOfNode,
					Metrics:      analyzer.(func(string) map[string]int)(path),
					Children:     make(map[string]string),
				}
			}

			if parent, exists := fileMetrics[parentPath]; exists && parentPath != relativePath {
				parent.Children[basename] = relativePath
				recursiveMetricAggregation(fileMetrics, parentPath, relativePath)
			} else {
				log.Printf("[DEBUG] No ancestor to %s ", relativePath)
			}

			return nil
		})
	if err != nil {
		log.Printf("[WARN] %s", err)
	}

	return fileMetrics
}

func recursiveMetricAggregation(fileMetrics map[string]core.FileData, parentPath string, relativePath string) {

	parent := fileMetrics[parentPath]
	for _, metricName := range aggregationMetrics {
		if _, exists := fileMetrics[relativePath].Metrics[metricName]; exists {
			childValue := fileMetrics[relativePath].Metrics[metricName]

			if _, exists := parent.Metrics[metricName]; exists {
				parent.Metrics[metricName] += childValue
			} else {
				parent.Metrics[metricName] = childValue
			}
		}
	}

	parentBasename := filepath.Base(parentPath)
	grandparentPath := strings.TrimSuffix(parentPath, string(filepath.Separator)+parentBasename)

	if grandparentPath == "" || parentPath == grandparentPath {
		grandparentPath = "."
	}

	if _, exists := fileMetrics[grandparentPath]; exists && parentPath != grandparentPath {
		recursiveMetricAggregation(fileMetrics, grandparentPath, relativePath)
	}
}
