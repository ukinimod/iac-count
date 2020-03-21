package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestLoc(t *testing.T) {
	tests := map[string]metricTest{
		"successful loc Calculation": {
			path:       "test/data/main.yml",
			content:    input.ReadFileToString("test/data/main.yml"),
			calculator: LocCalculator{},
			output:     27,
		},
	}

	runMetricTest(t, tests)
}
