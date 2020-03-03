package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestRloc(t *testing.T) {
	path := "test/data/main.yml"
	yamlString := input.ReadFileToString(path)
	got := RlocCalculator{}.Analyze(path, yamlString).Value()
	want := 26

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
