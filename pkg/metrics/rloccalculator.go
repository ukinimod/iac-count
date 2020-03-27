package metrics

import (
	"bufio"
	"regexp"
	"strings"
)

type RlocCalculator struct {
}

func (calculator RlocCalculator) Analyze(path, content string) Metric {
	re := regexp.MustCompile(`^\s*[^#\s]`)
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		text := scanner.Text()
		if re.FindStringIndex(text) != nil && text != "---" {
			count++
		}
	}

	return Rloc{
		Val: count,
	}
}
