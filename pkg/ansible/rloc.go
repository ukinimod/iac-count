package ansible

import (
	"bufio"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/MaibornWolff/iac-count/pkg/metrics"
)

type RlocCalculator struct {
}

func (calculator RlocCalculator) IsFileValidForMetric(path string) bool {
	return filepath.Ext(path) == ".yml" || filepath.Ext(path) == ".yaml"
}

func (calculator RlocCalculator) Analyze(path, content string) metrics.Metric {
	re := regexp.MustCompile(`^\s*[^#\s]`)
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		text := scanner.Text()
		if re.FindStringIndex(text) != nil && text != "---" {
			count++
		}
	}

	return metrics.Rloc{
		Val: count,
	}
}
