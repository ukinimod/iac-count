package ansible

import (
	"testing"
)

func TestFiles(t *testing.T) {
	tests := map[string]metricTest{
		"successful files Calculation": {
			path:       "test/data",
			content:    "",
			calculator: FilesCalculator{},
			output:     12,
		},
	}

	runMetricTest(t, tests)
}
