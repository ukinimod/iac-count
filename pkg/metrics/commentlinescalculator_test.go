package metrics

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestCommentLines(t *testing.T) {
	tests := map[string]metricTest{
		"successful comments Calculation": {
			path:       "test/data/main.yml",
			content:    input.ReadFileToString("test/data/main.yml"),
			calculator: CommentLinesCalculator{},
			output:     2,
		},
	}

	runMetricTest(t, tests)
}
