package core

import (
	"testing"
)

func TestTextfileAnalyze(t *testing.T) {
	tests := map[string]metricTest{
		"successful files Calculation": {
			path:           "test/data/taskfile.yaml",
			subjectCreator: TextfileCreator{},
			numberOfNodes:  1,
		},
	}

	runMetricTest(t, tests)
}
