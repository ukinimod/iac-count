package core

import (
	"testing"
)

func TestFileAnalyze(t *testing.T) {
	tests := map[string]metricTest{
		"successful files Calculation": {
			path:           "test/data/taskfile.yaml",
			subjectCreator: FileCreator{},
			numberOfNodes:  1,
		},
	}

	runMetricTest(t, tests)
}
