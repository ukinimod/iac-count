package ansible

import (
	"testing"

	input "github.com/MaibornWolff/iac-count/pkg/input"
)

func TestDependencies(t *testing.T) {
	yamlString := input.ReadFileToString("test/data/roles/example/meta/main.yml")
	got := DependencyCalculator{}.analyzeContent(yamlString).value()
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
