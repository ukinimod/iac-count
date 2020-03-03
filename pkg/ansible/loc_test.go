package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestLoc(t *testing.T) {
	path := "test/data/main.yml"
	yamlString := input.ReadFileToString(path)
	got := LocCalculator{}.Analyze(path, yamlString).Value()
	want := 28

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
