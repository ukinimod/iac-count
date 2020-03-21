package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursiveFileCount(t *testing.T) {
	tests := map[string]struct {
		path   string
		output int
	}{
		"existing file": {
			path:   "fileUtil.go",
			output: 1,
		},
		"invalid path": {
			path:   "nonexistingfile",
			output: 0,
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		output := RecursiveFileCount(test.path)
		assert.Equal(t, test.output, output)
	}
}
