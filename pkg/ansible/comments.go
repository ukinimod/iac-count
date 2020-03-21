package ansible

import (
	"bufio"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type CommentsCalculator struct {
}

func (calculator CommentsCalculator) isFileValidForMetric(path string) bool {
	return filepath.Ext(path) == ".yml"
}

func (calculator CommentsCalculator) Analyze(path, content string) metrics.Metric {
	if !calculator.isFileValidForMetric(path) {
		return nil
	}

	re := regexp.MustCompile(`^\s*#`)
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		text := scanner.Text()
		if re.FindStringIndex(text) != nil {
			count++
		}
	}

	return metrics.Comments{
		Val: count,
	}
}
