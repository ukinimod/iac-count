package core

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestDirectoryAnalyze(t *testing.T) {
	tests := map[string]metricTest{
		"successful files Calculation": {
			path:           "test/data",
			subjectCreator: DirectoryCreator{},
			numberOfNodes:  28,
		},
	}

	runMetricTest(t, tests)
}

type metricTest struct {
	path           string
	subjectCreator SubjectCreator
	numberOfNodes  int
}

func runMetricTest(t *testing.T, tests map[string]metricTest) {
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		subject := test.subjectCreator.CreateFromPath(test.path, nil)
		result := subject.Analyze()
		log.Infof("Result: %v", result)
		assert.Equal(t, test.numberOfNodes, len(result))
	}
}
