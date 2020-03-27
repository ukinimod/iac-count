package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadYamlAsList(t *testing.T) {
	tests := map[string]struct {
		content   string
		lenOutput int
	}{
		"valid yaml": {
			content:   ReadFileToString("test/data/taskfile.yaml"),
			lenOutput: 9,
		},
		"invalid yaml": {
			content:   " as as as ",
			lenOutput: 0,
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		output := len(ReadYamlAsList(test.content))
		assert.Equal(t, test.lenOutput, output)
	}
}

func TestReadYamlAsMap(t *testing.T) {
	exampleData := ReadFileToString("test/data/varsfile.yaml")
	got := ReadYamlAsMap(exampleData)["simple_var"]
	want := " ablas"

	assert.Equal(t, want, got)
}

func TestReadYamlAsMapOnFailure(t *testing.T) {
	invalidYaml := " as as as"
	got := len(ReadYamlAsMap(invalidYaml))
	want := 0

	assert.Equal(t, want, got)
}
