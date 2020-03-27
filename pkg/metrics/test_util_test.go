package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type metricTest struct {
	path       string
	content    string
	calculator MetricCalculator
	output     int
}

func runMetricTest(t *testing.T, tests map[string]metricTest) {
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		result := test.calculator.Analyze(test.path, test.content).Value()
		assert.Equal(t, test.output, result)
	}
}
