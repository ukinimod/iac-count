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
		"directory with multiple files": {
			path:   ".",
			output: 3,
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

func TestIsTextFile(t *testing.T) {
	tests := map[string]struct {
		path   string
		result bool
	}{
		"json": {
			path:   "test/data/main.json",
			result: true,
		},
		"yml": {
			path:   "test/data/main.yml",
			result: true,
		},
		"txt": {
			path:   "test/data/main.txt",
			result: true,
		},
		"mp3": {
			path:   "test/data/main.mp3",
			result: false,
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		result := IsTextFile(test.path)
		assert.Equal(t, test.result, result)
	}
}
