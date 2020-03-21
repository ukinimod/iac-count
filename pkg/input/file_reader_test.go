package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToString(t *testing.T) {
	tests := map[string]struct {
		path      string
		lenOutput int
	}{
		"existing file": {
			path:      "test/data/taskfile.yaml",
			lenOutput: 852,
		},
		"invalid path": {
			path:      "test/data/unknownfile",
			lenOutput: 0,
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		output := len(ReadFileToString(test.path))
		assert.Equal(t, test.lenOutput, output)
	}
}
