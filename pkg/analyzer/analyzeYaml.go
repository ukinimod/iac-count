package analyzer

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/ukinimod/iac-count/pkg/core"
	input "github.com/ukinimod/iac-count/pkg/input"
)

func getCountLineMatchingFunc(pattern string) func(int, string) int {
	re := regexp.MustCompile(pattern)
	f := func(oldMetric int, text string) int {
		if re.FindStringIndex(text) != nil {
			return oldMetric + 1
		}
		return oldMetric
	}

	return f
}

func getMaximalIndentationFunction() func(int, string) int {
	const indentation = 2
	re := regexp.MustCompile(`^\s*\w+:.*`)
	f := func(oldMetric int, text string) int {
		if re.FindStringIndex(text) != nil {
			unindentedText := strings.TrimLeftFunc(text, func(s rune) bool { return s == ' ' })
			v := (len(text) - len(unindentedText)) / indentation
			if v > oldMetric {
				return v
			}
		}
		return oldMetric
	}

	return f
}

func initMetricFuncs() map[string]func(int, string) int {
	metricFuncs := make(map[string]func(int, string) int)

	metricFuncs[core.Rloc] = getCountLineMatchingFunc(`^\s*[^#].*`)
	metricFuncs[core.CommentLines] = getCountLineMatchingFunc(`^\s*#`)
	metricFuncs[core.MaxNL] = getMaximalIndentationFunction()

	return metricFuncs
}

func calcMetricPerLine(sourceString string, metricFuncs map[string]func(int, string) int) map[string]int {
	metrics := make(map[string]int, len(metricFuncs))
	scanner := bufio.NewScanner(strings.NewReader(sourceString))
	for scanner.Scan() {
		text := scanner.Text()
		for key := range metricFuncs {
			metrics[key] = metricFuncs[key](metrics[key], text)
		}
	}
	return metrics
}

func analyzeYamlString(content string) map[string]int {
	metrics := calcMetricPerLine(content, initMetricFuncs())

	metrics[core.Loc] = numberOfLines(content)

	metrics[core.Complexity] = 1
	metrics[core.Files] = 1

	return metrics
}

func analyzeYaml(path string) map[string]int {
	content := input.ReadFileToString(path)

	return analyzeYamlString(content)
}
