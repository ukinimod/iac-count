package metrics

import (
	"bufio"
	"regexp"
	"strings"
)

type CommentlinesCalculator struct {
}

func (calculator CommentlinesCalculator) Analyze(path, content string) Metric {
	re := regexp.MustCompile(`^\s*#`)
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		text := scanner.Text()
		if re.FindStringIndex(text) != nil {
			count++
		}
	}

	return Commentlines{
		Val: count,
	}
}
