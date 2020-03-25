package core

import (
	"testing"
)

func TestYamlfileAnalyze(t *testing.T) {
	tests := map[string]metricTest{
		"successful files Calculation": {
			path:           "test/data/taskfile.yaml",
			subjectCreator: YamlfileCreator{},
			numberOfNodes:  1,
		},
	}

	runMetricTest(t, tests)
}
